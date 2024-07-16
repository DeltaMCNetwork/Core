package server

type IProtocolHandler interface {
	HandlePacket(int32, IBuffer, IConnection, *MinecraftServer)
	HandlePacketStatus(int32, IBuffer, IConnection, *MinecraftServer)
	HandlePacketPlay(int32, IBuffer, IConnection, *MinecraftServer)
	HandlePacketLogin(int32, IBuffer, IConnection, *MinecraftServer)
	HandlePacketPing(int32, IBuffer, IConnection, *MinecraftServer)
}

type ProtocolHandler struct {
	IProtocolHandler
}

func createBasicProtocolHandler() IProtocolHandler {
	return &ProtocolHandler{}
}

func (handler *ProtocolHandler) HandlePacket(packetId int32, buffer IBuffer, conn IConnection, server *MinecraftServer) {
	Debug("Packet mode is %d", conn.GetPacketMode())
	switch conn.GetPacketMode() {
	case PacketModeLogin:
		handler.HandlePacketLogin(packetId, buffer, conn, server)
	case PacketModeStatus:
		handler.HandlePacketStatus(packetId, buffer, conn, server)
	case PacketModePing:
		handler.HandlePacketPing(packetId, buffer, conn, server)
	case PacketModePlay:
		handler.HandlePacketPlay(packetId, buffer, conn, server)
	}
}

func (handler *ProtocolHandler) HandlePacketLogin(packetId int32, buffer IBuffer, conn IConnection, server *MinecraftServer) {
	switch packetId {
	case 0:
		username := buffer.ReadString()

		Info(username + " Joined your server")
	}
}

func (handler *ProtocolHandler) HandlePacketStatus(packetId int32, buffer IBuffer, conn IConnection, server *MinecraftServer) {

}

func (handler *ProtocolHandler) HandlePacketPing(packetId int32, buffer IBuffer, conn IConnection, server *MinecraftServer) {
	switch packetId {
	case 0:
		protocolVersion := buffer.ReadVarInt()
		serverAddress := buffer.ReadString()
		port := buffer.ReadUInt16()
		nextState := buffer.ReadVarInt()

		if protocolVersion != 47 {
			Info("Incorrect version!")
		}

		Debug("Address: %s Port: %d State: %d", serverAddress, port, nextState)
		Debug("Client Protocol is Version %d", protocolVersion)

		conn.SetPacketMode(int(nextState))
	}
}

func (handler *ProtocolHandler) HandlePacketPlay(packetId int32, buffer IBuffer, conn IConnection, server *MinecraftServer) {

}
