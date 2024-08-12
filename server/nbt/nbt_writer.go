package nbt

import "net/deltamc/server"

func Write(tag NbtCompound, buf server.IBuffer) {
	for key, value := range tag.data {
		//WriteString(key, buf)
		WriteValueTag(value, buf, key, true)
	}
}

func WriteValue(value any, buf server.IBuffer) {
	WriteValueTag(value, buf, "", true)
}

func WriteString(value string, buf server.IBuffer) {
	buf.WriteInt16(int16(len(value)))
	buf.Write([]byte(value))
}

func WriteValueTag(value any, buf server.IBuffer, name string, writeTag bool) {
	switch val := value.(type) {
	case byte:
		buf.WriteByte(TagByte)
		buf.WriteByte(val)
	case int16:
		if writeTag {
			buf.WriteByte(TagShort)
		}
		buf.WriteInt16(val)
	case uint16:
		if writeTag {
			buf.WriteByte(TagShort)
		}
		buf.WriteUInt16(val)
	case int32:
		if writeTag {
			buf.WriteByte(TagInt)
		}
		buf.WriteInt(val)
	case uint32:
		if writeTag {
			buf.WriteByte(TagInt)
		}
		buf.WriteUInt(val)
	case int64:
		if writeTag {
			buf.WriteByte(TagLong)
		}
		buf.WriteLong(val)
	case uint64:
		if writeTag {
			buf.WriteByte(TagLong)
		}
		buf.WriteULong(val)
	case float32:
		if writeTag {
			buf.WriteByte(TagFloat)
		}
		buf.WriteFloat(val)
	case float64:
		if writeTag {
			buf.WriteByte(TagFloat)
		}
		buf.WriteDouble(val)
	case []byte:
		if writeTag {
			buf.WriteByte(TagByteArray)
		}
		buf.WriteInt(int32(len(val)))
		buf.Write(val)
	case string:
		if writeTag {
			buf.WriteByte(TagString)
		}
		buf.WriteInt16(int16(len(val)))
		buf.Write([]byte(val))
	case []int32:
		if writeTag {
			buf.WriteByte(TagIntArray)
		}

		buf.WriteInt(int32(len(val)))

		for in := range val {
			buf.WriteInt(val[in])
		}

		return
	case []int64:
		if writeTag {
			buf.WriteByte(TagLongArray)
		}

		buf.WriteInt(int32(len(val)))

		for in := range val {
			buf.WriteLong(val[in])
		}

		return
	case []any:
		if writeTag {
			buf.WriteByte(TagList)
		}
		buf.WriteInt(int32(len(val)))

		for in := range val {
			WriteValueTag(val[in], buf, "", false)
		}
	case NbtCompound:
		if writeTag {
			buf.WriteByte(TagCompound)
		}

		Write(val, buf)
	}
}
