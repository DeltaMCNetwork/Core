package server

import (
	"net/deltamc/server/component"
	"net/deltamc/server/status"

	uuid "github.com/satori/go.uuid"
)

type ProtocolFunc = func(IBuffer, IConnection, *MinecraftServer) bool

type ProtocolTable struct {
	funcs map[int32]ProtocolFunc
	index int32
}

func CreateProtocolTable() *ProtocolTable {
	table := &ProtocolTable{
		index: 0,
		funcs: make(map[int32]ProtocolFunc, 0),
	}

	initTable(table)

	return table
}

func initTable(table *ProtocolTable) {
	table.IotaRegister(func(buffer IBuffer, conn IConnection, server *MinecraftServer) bool {
		keepAlive := &ClientKeepAlive{}
		keepAlive.Read(buffer)

		server.packetHandler.HandleKeepAlive(*keepAlive, conn.GetPlayer())

		return true
	})
}

func (table *ProtocolTable) HandlePacket(id int32, buffer IBuffer, conn IConnection, server *MinecraftServer) {
	switch conn.GetPacketMode() {
	case PacketModePlay:
		if !table.funcs[id](buffer, conn, server) {
			conn.GetPlayer().Disconnect(component.NewTextComponent("ma friend"))
		}
	case PacketModeLogin:
		switch id {
		case ClientLoginStartPacket: // Login Start
			name := buffer.ReadString()
			player := server.playerCreate(name)
			player.SetConnection(conn)
			conn.SetPlayer(player)

			if server.GetAuthenticator() != nil {
				// online mode enabled
				conn.SendPacket(CreateServerEncryptionRequest(server.GetKeypair().Public, GenerateVerificationToken()))
			} else {
				player.SetUuid(uuid.FromStringOrNil("Offline:" + player.GetUsername()))
				player.SetAuthenticated(true)
				completeLogin(player)
			}
		case ClientEncryptionResponsePacket: // Encryption Request
			sharedSecret := buffer.ReadByteArray()
			verifyToken := buffer.ReadByteArray()
			decrypted, err := server.GetKeypair().Decrypt(sharedSecret)

			if err != nil {
				Error("Error decrypting sharedSecret: %s", err.Error())
			}

			sharedSecret = decrypted
			// USE THIS
			decVerifyToken, err := server.GetKeypair().Decrypt(verifyToken)

			if err != nil {
				Error("Error decrypting verify token: %s", err.Error())
			}

			verifyToken = decVerifyToken

			if verifyToken == nil {
				panic("fine")
			}

			conn.EnableEncryption(sharedSecret)

			//TODO: where to put verify token
			Info("Encryption response")

			player := conn.GetPlayer()
			authResult := server.GetAuthenticator().Authenticate(player, server, sharedSecret)
			switch authResult.Result {
			case AuthSuccess:
				completeLogin(player)
				player.SetAuthenticated(true)
			case AuthFail:
				player.Disconnect(component.NewTextComponent("buy mineccraft poor ass").WithColor(component.Red))
			case AuthError:
				player.Disconnect(component.NewTextComponent("Something went wrong during authentication. Please try again later.").WithColor(component.Red))
			}
		}
	case PacketModeStatus:
		switch id {
		case 0x00: // status request
			event := &ServerListPingEvent{}

			//TODO: player count
			//TODO: player samples

			event.Response = status.NewResponse().
				WithVersion(PROTOCOL_VERSION, PROTOCOL_NAME).
				WithInfo(0, 20).
				WithSamples([]*status.PlayerSample{status.NewPlayerSample("Dream", "474ee1bc-e3e1-4672-9b37-284a6993d9b7")}).
				WithDescription(component.NewTextComponent("Welcome to the Deltamc Server"))

			if favicon := server.GetFavicon(); favicon != "" {
				event.Response = event.Response.WithFavicon(favicon)
			}

			server.GetInjectionManager().Post(event)

			packet := CreateServerStatusResponse(event.Response)
			conn.SendPacket(packet)
		case 0x01: // status ping
			payload := buffer.ReadLong()
			conn.SendPacket(CreateServerStatusPong(payload))
		}
	case PacketModeHandshake:
		packet := NewClientHandshake()
		packet.Read(buffer)

		if packet.NextState > 0 && packet.NextState < 3 {
			// check if it's a valid protocol version
			conn.SetPacketMode(int(packet.NextState))
		} else {
			conn.Remove()
		}
	}
}

func (table *ProtocolTable) IotaRegister(f ProtocolFunc) {
	table.funcs[table.index] = f
	table.index++
}

func (table *ProtocolTable) Register(index int32, f ProtocolFunc) {
	table.funcs[index] = f
}

func completeLogin(player IPlayer) {
	Info("complete login")
	player.SendPacket(CreateServerLoginSuccess(player.GetUuid().String(), player.GetUsername()))
}
