package server

import (
	"fmt"
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
			conn.GetPlayer().Disconnect("bro ur bad")
		}
	case PacketModeLogin:
		fmt.Println("holy shit login")
		switch id {
		case 0x00: // Login Start
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
		case 0x01: // Encryption Request
			sharedSecret := buffer.ReadByteArray()
			verifyToken := buffer.ReadByteArray()

			server.GetKeypair().Decrypt(sharedSecret)
			server.GetKeypair().Decrypt(verifyToken)

			conn.EnableEncryption(sharedSecret)

			//TODO: where to put verify token

			go func() {
				player := conn.GetPlayer()
				authResult := server.GetAuthenticator().Authenticate(player, server)
				switch authResult.Result {
				case AuthSuccess:
					completeLogin(player)
				case AuthFail:
					player.Disconnect("buy minecraft poor ass")
				case AuthError:
					player.Disconnect("Something went wrong while authentication you")
				}
			}()
		}
	case PacketModeStatus:
		break
	case PacketModeHandshake:
		packet := NewClientHandshake()
		packet.Read(buffer)

		if packet.NextState > 0 && packet.NextState < 3 {
			// check if it's a valid protocol version
			conn.SetPacketMode(packet.NextState)
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
	player.SendPacket(CreateServerLoginSuccess(player.GetUuid().String(), player.GetUsername()))
}
