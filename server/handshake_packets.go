package server

type ClientHandshake struct {
	ProtocolVersion int
	ServerAddress   string
	ServerPort      uint16
	NextState       byte
}

func NewClientHandshake() *ClientHandshake {
	return &ClientHandshake{}
}

func (handshake *ClientHandshake) Read(buf IBuffer) {
	handshake.ProtocolVersion = int(buf.ReadVarInt())
	handshake.ServerAddress = buf.ReadString()
	handshake.ServerPort = buf.ReadUInt16()
	handshake.NextState = buf.ReadByte()
}
