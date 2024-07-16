package server

type IPacket interface {
	GetPacketId() int
	Read(*IBuffer) IPacket
	Write(*IBuffer) *IBuffer
}
