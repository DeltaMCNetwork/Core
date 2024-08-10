package server

import (
	"net/deltamc/server/component"
	"strings"
)

type IPlayer interface {
	GetUsername() string
	SetUsername(string)
	GetUuid() UUID
	SetUuid(UUID)
	GetConnection() IConnection
	SetConnection(IConnection)
	GetIP() string

	Disconnect(*component.TextComponent)
	SendPacket(ServerPacket)
}

type BasicPlayer struct {
	IPlayer
	username   string
	uuid       UUID
	connection IConnection
}

func createBasicPlayer(username string) IPlayer {
	return &BasicPlayer{
		username: username,
	}
}

func (player *BasicPlayer) GetIP() string {
	ipStr := player.GetConnection().GetIP()
	indexOf := strings.Index(ipStr, ":")

	return ipStr[:indexOf]
}

func (player *BasicPlayer) GetUsername() string {
	return player.username
}

func (player *BasicPlayer) SetUsername(value string) {
	player.username = value
}

func (player *BasicPlayer) GetUuid() UUID {
	return player.uuid
}

func (player *BasicPlayer) SetUuid(uuid UUID) {
	player.uuid = uuid
}

func (player *BasicPlayer) GetConnection() IConnection {
	return player.connection
}

func (player *BasicPlayer) SetConnection(conn IConnection) {
	player.connection = conn
}

func (player *BasicPlayer) Disconnect(text *component.TextComponent) {
	player.SendPacket(CreateServerDisconnect(text))

	player.connection.Remove()
}

func (player *BasicPlayer) SendPacket(packet ServerPacket) {
	player.connection.SendPacket(packet)
}
