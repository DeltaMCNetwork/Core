package server

type IProtocolHandler interface {
	HandlePacket(IBuffer, IConnection, MinecraftServer)
	HandlePacketStatus(IBuffer, IConnection, MinecraftServer)
	HandlePacketPlay(IBuffer, IConnection, MinecraftServer)
	HandlePacketLogin(IBuffer, IConnection, MinecraftServer)
	HandlePacketPing(IBuffer, IConnection, MinecraftServer)
}

type ProtocolHandler struct {
	IProtocolHandler
}

func (handler *ProtocolHandler) HandlePacket(buffer IBuffer, conn IConnection, server MinecraftServer) {
	switch conn.GetPacketMode() {
	case PacketModeLogin:
		handler.HandlePacketLogin(buffer, conn, server)
	case PacketModeStatus:
		handler.HandlePacketStatus(buffer, conn, server)
	case PacketModePing:
		handler.HandlePacketPing(buffer, conn, server)
	case PacketModePlay:
		handler.HandlePacketPlay(buffer, conn, server)
	}
}

func (handler *ProtocolHandler) HandlePacketLogin(buffer IBuffer, conn IConnection, server MinecraftServer) {

}

func (handler *ProtocolHandler) HandlePacketStatus(buffer IBuffer, conn IConnection, server MinecraftServer) {

}

func (handler *ProtocolHandler) HandlePacketPing(buffer IBuffer, conn IConnection, server MinecraftServer) {

}

func (handler *ProtocolHandler) HandlePacketPlay(buffer IBuffer, conn IConnection, server MinecraftServer) {

}
