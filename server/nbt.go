package server

import (
	"bytes"
	"compress/gzip"
	"io"
)

// this is a glorified custom type, yippe
type NbtCompound map[string]any

func (compound NbtCompound) Get(key string) any {
	return compound[key]
}

func (compound NbtCompound) Set(key string, value any) {
	compound[key] = value
}

type NbtList []any

func (list NbtList) Length() int {
	return len(list)
}

type NbtByteArray []uint8

func (list NbtByteArray) Length() int {
	return len(list)
}

type NbtIntArray []int32

func (list NbtIntArray) Length() int {
	return len(list)
}

type NbtLongArray []int64

func (list NbtLongArray) Length() int {
	return len(list)
}

type NbtTagType = byte

const (
	TagEnd NbtTagType = iota
	TagByte
	TagShort
	TagInt
	TagLong
	TagFloat
	TagDouble
	TagByteArray
	TagString
	TagList
	TagCompound
	TagIntArray
	TagLongArray
)

func NbtGetType(value any) NbtTagType {
	switch value := value.(type) {
	case int:
		switch {
		case value < 127 && value > -128:
			return TagByte
		case value < 32767 && value > -32768:
			return TagShort
		case value < 2147483647 && value > -2147483648:
			return TagInt
		case value < 9223372036854775807 && value > -9223372036854775807:
			return TagLong
		default:
			return TagEnd
		}
	case int8, uint8, bool:
		return TagByte
	case int16, uint16:
		return TagShort
	case int32, uint32:
		return TagInt
	case int64, uint64:
		return TagLong
	case float32:
		return TagFloat
	case float64:
		return TagDouble
	case []byte, NbtByteArray:
		return TagByteArray
	case []int64, []uint64, NbtLongArray:
		return TagLongArray
	case []int32, []uint32, NbtIntArray:
		return TagIntArray
	case NbtCompound, map[string]any:
		return TagCompound
	case []any, NbtList:
		return TagList
	case string:
		return TagString
	default:
		return TagEnd
	}
}

func NbtReadGzip(data []byte, buf IBuffer) *NbtCompound {
	reader := bytes.NewReader(data)
	gzReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil
	}

	dat, err := io.ReadAll(gzReader)
	if err != nil {
		return nil
	}
	buf.SetData(dat)

	return NbtRead(buf)
}

func NbtRead(buf IBuffer) *NbtCompound {
	compound := &NbtCompound{}
	tag := buf.ReadByte()

	if tag != TagCompound {
		return compound
	}

	Info("Len of buffer is %d", buf.GetLength())
	_ = readString(buf)
	readWithName(compound, buf)

	return compound
}

func readString(buf IBuffer) string {
	return string(buf.Read(int(buf.ReadInt16())))
}

func writeString(buf IBuffer, value string) {
	buf.WriteInt16(int16(len(value)))
	buf.Write([]byte(value))
}

func readWithName(compound *NbtCompound, buf IBuffer) {
	for {
		var tag NbtTagType = buf.ReadByte()
		if tag == TagEnd {
			return
		}

		name := readString(buf)
		val := readValue(buf, tag)
		if val == nil {
			Info("Nil in read nbt in tag %d", tag)
		}

		(*compound)[name] = val
	}
}

func readValue(buf IBuffer, tag NbtTagType) any {
	switch tag {
	case TagByte:
		return buf.ReadByte()
	case TagShort:
		return buf.ReadInt16()
	case TagInt:
		return buf.ReadInt()
	case TagLong:
		return buf.ReadLong()
	case TagFloat:
		return buf.ReadFloat()
	case TagDouble:
		return buf.ReadDouble()
	case TagByteArray:
		return buf.Read(int(buf.ReadInt()))
	case TagString:
		return readString(buf)
	case TagList:
		var listType NbtTagType = buf.ReadByte()
		len := int(buf.ReadInt())
		data := make(NbtList, len)
		for i := 0; i < len; i++ {
			data[i] = readValue(buf, listType)
		}

		return data
	case TagCompound:
		data := &NbtCompound{}
		readWithName(data, buf)

		return *data
	case TagLongArray:
		len := int(buf.ReadInt())
		data := make(NbtLongArray, len)
		for i := 0; i < len; i++ {
			data[i] = buf.ReadLong()
		}

		return data
	case TagIntArray:
		len := int(buf.ReadInt())
		data := make(NbtIntArray, len)
		for i := 0; i < len; i++ {
			data[i] = buf.ReadInt()
		}

		return data
	}

	return nil
}

//write

func NbtWrite(buf IBuffer, compound NbtCompound) {
	buf.WriteByte(TagCompound)
	writeString(buf, "")
	writeWithName(buf, compound)
}

func NbtWriteGzip(buf IBuffer, compound NbtCompound) []byte {
	NbtWrite(buf, compound)

	newBuf := new(bytes.Buffer)
	gWriter := gzip.NewWriter(newBuf)
	gWriter.Write(buf.GetBytes())
	gWriter.Close()

	return newBuf.Bytes()
}

func writeWithName(buf IBuffer, compound NbtCompound) {
	for k, v := range compound {
		tag := NbtGetType(v)

		if tag == TagEnd {
			panic("why tf is this end tag key is " + k)
		}

		buf.WriteByte(tag)
		writeString(buf, k)
		writeValue(buf, tag, v)
	}
	buf.WriteByte(TagEnd)
}

func writeValue(buf IBuffer, tag NbtTagType, value any) {
	switch tag {
	case TagByte:
		switch v := value.(type) {
		case bool:
			buf.WriteBool(v)
		case uint8:
			buf.WriteUInt8(v)
		case int8:
			buf.WriteInt8(v)
		}
	case TagCompound:
		writeWithName(buf, (value.(NbtCompound)))
	case TagShort:
		buf.WriteInt16(value.(int16))
	case TagInt:
		switch val := value.(type) {
		case int:
			buf.WriteInt(int32(val))
		}
		buf.WriteInt(value.(int32))
	case TagLong:
		buf.WriteLong(value.(int64))
	case TagFloat:
		buf.WriteFloat(value.(float32))
	case TagDouble:
		buf.WriteDouble(value.(Double))
	case TagString:
		writeString(buf, value.(string))
	case TagIntArray:
		arr := value.(NbtIntArray)
		len := len(arr)
		buf.WriteInt(int32(len))
		for i := 0; i < len; i++ {
			buf.WriteInt(arr[i])
		}
	case TagLongArray:
		arr := value.(NbtLongArray)
		len := len(arr)
		buf.WriteInt(int32(len))
		for i := 0; i < len; i++ {
			buf.WriteLong(arr[i])
		}
	case TagByteArray:
		arr := value.([]uint8)
		len := len(arr)
		buf.WriteInt(int32(len))
		for i := 0; i < len; i++ {
			buf.WriteByte(arr[i])
		}
	case TagList:
		arr := value.(NbtList)
		len := len(arr)
		if len == 0 {
			buf.Write([]byte{0, 0, 0, 0, 0})
			return
		}
		tag := NbtGetType(arr[0])
		buf.WriteByte(tag)
		buf.WriteInt(int32(len))
		for i := range arr {
			writeValue(buf, tag, arr[i])
		}
	default:
		Info("Wrote nothing what tehs igma")
	}
}
