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
	ReadUUID() uuid.UUID
}
