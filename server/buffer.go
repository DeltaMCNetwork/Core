package server

type IBuffer interface {
	Read(count int) []byte
	ReadUInt8() uint8
	ReadInt8() int8
}
