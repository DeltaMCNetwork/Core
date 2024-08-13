package nbt

import "net/deltamc/server"

var nbtWriters = make(map[NbtTagType]func(*NbtCompound, string, server.IBuffer, any), 0)

func InitWriter() {
	nbtWriters[TagByte] = func(nc *NbtCompound, s string, i server.IBuffer, value any) {
		i.WriteByte(value.(byte))
	}
	nbtWriters[TagByteArray] = func(nc *NbtCompound, s string, i server.IBuffer, a any) {
		i
	}
}

func writeString(name string, buf server.IBuffer) {
	buf.WriteInt16(int16(len(name)))
	buf.Write([]byte(name))
}

func Write(compound *NbtCompound, buf server.IBuffer) {
	buf.WriteByte(TagCompound)
	writeString(compound.name, buf)

	for k, v := range compound.data {
		tag := GetTagType(v)

		switch tag {
		case TagEnd:
			panic("Impossible NBT encoding")
		default:
			buf.WriteByte(tag)
			writeString(k, buf)
			nbtWriters[tag](compound, k, buf, v)
		}
	}

	buf.WriteByte(TagEnd)
}

func GetTagType(value any) NbtTagType {
	switch value.(type) {
	case byte:
		return TagByte
	case []byte:
		return TagByteArray
	case []int32:
		return TagIntArray
	case []int64:
		return TagLongArray
	case NbtCompound:
		return TagCompound
	case float64:
		return TagDouble
	case float32:
		return TagFloat
	case int32:
		return TagInt
	case int64:
		return TagLong
	case string:
		return TagString
	case int16:
		return TagShort
	case bool:
		return TagByte
	case []NbtCompound, []int16, []float32, []float64:
		return TagList
	default:
		return TagEnd
	}
}
