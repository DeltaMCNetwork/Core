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
				// online mode enbabled
				conn.SendPacket(CreateServerEncryptionRequest(server.GetKeypair().Public, GenerateVerificationToken()))
			} else {
				player.SetUuid(uuid.FromStringOrNil("Offline:" + player.GetUsername()))
				completeLogin(player)
			}
		case ClientEncryptionResponsePacket: // Encryption Request
			sharedSecret := buffer.ReadByteArray()
			Info("Length of sharedSecret is %d", len(sharedSecret))
			verifyToken := buffer.ReadByteArray()
			decrypted, err := server.GetKeypair().Decrypt(sharedSecret)

			if err != nil {
				Error("Error decrypting sharedSecret: %s", err.Error())
			}

			sharedSecret = decrypted
			decVerifyToken, err := server.GetKeypair().Decrypt(verifyToken)

			if err != nil {
				Error("Error decrypting verify token: %s", err.Error())
			}

			verifyToken = decVerifyToken
			conn.EnableEncryption(sharedSecret)

			//TODO: where to put verify token
			Info("Encryption response")

			player := conn.GetPlayer()
			authResult := server.GetAuthenticator().Authenticate(player, server, sharedSecret)
			switch authResult.Result {
			case AuthSuccess:
				completeLogin(player)
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
				WithDescription(component.NewTextComponent("Welcome to the Deltamc Server")).
				WithFavicon("data:image/png;base64,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")

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
