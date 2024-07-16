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
	//Debug("Packet mode is %d", conn.GetPacketMode())
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

		player := server.playerCreate(username)
		player.SetConnection(conn)
		conn.SetPlayer(player)

		if len(username) == 0 {
			player.Disconnect("Invalid username!")

			return
		}

		if conn.GetProtocolVersion() != PROTOCOL_VERSION {
			player.Disconnect(Stringify("&6deltamc.net\n&c\n&7Invalid Minecraft version!\n&8Expected %d and got %d", PROTOCOL_VERSION, conn.GetProtocolVersion()))

			return
		}

		if USE_PROXY {
			uuid := buffer.ReadUUID()

			player.SetUuid(uuid)
		} else {
			player.SetUuid(CreateUUID())
		}
		// buffer

		Info(player.GetUsername() + " Joined your server")
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

		Debug("Address: %s Port: %d State: %d", serverAddress, port, nextState)
		Debug("Client Protocol is Version %d", protocolVersion)

		conn.SetProtocolVersion(int(protocolVersion))
		conn.SetPacketMode(int(nextState))
	}
}

func (handler *ProtocolHandler) HandlePacketPlay(packetId int32, buffer IBuffer, conn IConnection, server *MinecraftServer) {

}
