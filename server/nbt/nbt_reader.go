package nbt

import "net/deltamc/server"

var nbtReaders = make(map[NbtTagType]func(*NbtCompound, server.IBuffer), 0)

func InitReader() {
	nbtReaders[TagByte] = func(compound *NbtCompound, buf server.IBuffer) {
		name := readString(buf)
		value := buf.ReadByte()

		compound.Set(name, value)
	}
	nbtReaders[TagByteArray] = func(compound *NbtCompound, buf server.IBuffer) {
		name := readString(buf)
		len := buf.ReadInt()

		data := buf.Read(int(len))

		compound.Set(name, data)
	}
	nbtReaders[TagCompound] = func(compound *NbtCompound, buf server.IBuffer) {
		name := readString(buf)

		compound.Set(name, *Read(buf, name))
	}
	nbtReaders[TagDouble] = func(compound *NbtCompound, buf server.IBuffer) {
		name := readString(buf)
		value := buf.ReadDouble()

		compound.Set(name, value)
	}
	nbtReaders[TagFloat] = func(compound *NbtCompound, buf server.IBuffer) {
		name := readString(buf)
		value := buf.ReadFloat()

		compound.Set(name, value)
	}
	nbtReaders[TagInt] = func(compound *NbtCompound, buf server.IBuffer) {
		name := readString(buf)
		value := buf.ReadInt()

		compound.Set(name, value)
	}
	nbtReaders[TagIntArray] = func(compound *NbtCompound, buf server.IBuffer) {
		name := readString(buf)
		len := buf.ReadInt()

		data := make([]int, len)

		for i := 0; i < int(len); i++ {
			data[i] = int(buf.ReadInt())
		}

		compound.Set(name, data)
	}
	nbtReaders[TagList] = func(compound *NbtCompound, buf server.IBuffer) {
		name := readString(buf)
		listType := NbtTagType(buf.ReadByte())
		len := buf.ReadInt()

		switch listType {
		case TagCompound:
			data := make([]NbtCompound, len)
			for k := range data {
				data[k] = *Read(buf, readString(buf))
			}
			compound.Set(name, data)
		case TagByte:
			compound.Set(name, buf.Read(int(len)))
		case TagShort:
			data := make([]int16, len)
			for k := range data {
				data[k] = buf.ReadInt16()
			}
			compound.Set(name, data)
		case TagInt:
			data := make([]int32, len)
			for k := range data {
				data[k] = buf.ReadInt()
			}
			compound.Set(name, data)
		case TagLong:
			data := make([]int64, len)
			for k := range data {
				data[k] = buf.ReadLong()
			}
			compound.Set(name, data)
		case TagFloat:
			data := make([]float32, len)
			for k := range data {
				data[k] = buf.ReadFloat()
			}
			compound.Set(name, data)
		case TagDouble:
			data := make([]float64, len)
			for k := range data {
				data[k] = buf.ReadDouble()
			}
			compound.Set(name, data)
		default:
			panic("Invalid list type! wtf crash")
		}
	}
	nbtReaders[TagLong] = func(compound *NbtCompound, buf server.IBuffer) {
		name := readString(buf)
		value := buf.ReadLong()

		compound.Set(name, value)
	}
	nbtReaders[TagLongArray] = func(compound *NbtCompound, buf server.IBuffer) {
		name := readString(buf)
		len := buf.ReadInt()

		data := make([]int64, len)
		for k := range data {
			data[k] = buf.ReadLong()
		}

		compound.Set(name, data)
	}
	nbtReaders[TagShort] = func(compound *NbtCompound, buf server.IBuffer) {
		name := readString(buf)
		value := buf.ReadInt16()

		compound.Set(name, value)
	}
	nbtReaders[TagString] = func(compound *NbtCompound, buf server.IBuffer) {
		name := readString(buf)
		value := readString(buf)

		compound.Set(name, value)
	}
}

func readString(buf server.IBuffer) string {
	len := buf.ReadInt16()
	str := string(buf.Read(int(len)))

	return str
}

func Read(buf server.IBuffer, name string) *NbtCompound {
	compound := CreateCompound(name)

	for {
		tag := NbtTagType(buf.ReadByte())

		switch tag {
		case TagEnd:
			return compound
		default:
			nbtReaders[tag](compound, buf)
		}
	}
}
