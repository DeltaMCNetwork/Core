package server

type IBuffer interface {
	Read(count int32) []byte
	ReadUInt8() uint8
	ReadInt8() int8
	ReadUInt16() uint16
	ReadInt16() int16
	ReadInt() int32
	ReadUInt() uint32
	ReadVarInt() int32
	ReadFloat() float32
	ReadDouble() double
	ReadLong() int64
	ReadULong() uint64
	ReadPosition() *Position
	ReadString() string
	ReadUUID() UUID

	Write([]byte)
	WriteUInt8(uint8)
	WriteInt8(int8)
	WriteUInt16(uint16)
	WriteInt16(int16)
	WriteInt(int32)
	WriteUInt(uint32)
	WriteVarInt(int32)
	WriteFloat(float32)
	WriteDouble(double)
	WriteLong(int64)
	WriteULong(uint64)
	WritePosition(*Position)
	WriteString() string
	WriteUUID(UUID)

	GetBytes() []byte
	GetLength() int32
	GetPointer() int32
}

func createBasicBuffer() IBuffer {
	return &BasicBuffer{}
}

type BasicBuffer struct {
	IBuffer
	data    []byte
	pointer int32
}

func (buffer *BasicBuffer) Read(count int32) []byte {
	data := buffer.data[buffer.pointer : buffer.pointer+count]
	buffer.pointer += count

	return data
}

func (buffer *BasicBuffer) ReadUInt8() uint8 {
	return uint8(buffer.Read(1)[0])
}

func (buffer *BasicBuffer) ReadInt8() int8 {
	return int8(buffer.Read(1)[0])
}
