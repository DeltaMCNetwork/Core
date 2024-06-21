package server

type IBuffer interface {
	Read(count int) []byte
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
}
