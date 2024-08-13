package server

import (
	"net/deltamc/server/component"
)

type IPlayer interface {
	GetEntityId() int32
	GetUsername() string
	SetUsername(string)
	GetUuid() UUID
	SetUuid(UUID)
	GetConnection() IConnection
	SetConnection(IConnection)
	IsAuthenticated() bool
	SetAuthenticated(bool)
	GetIP() string

	SendMessage(*component.TextComponent)
	Disconnect(*component.TextComponent)
	SendPacket(ServerPacket)
}

type BasicPlayer struct {
	entityId      int32
	username      string
	uuid          UUID
	connection    IConnection
	authenticated bool
}

func createBasicPlayer(username string) IPlayer {
	return &BasicPlayer{
		username: username,
	}
}

func (player *BasicPlayer) GetEntityId() int32 {
	return player.entityId
}

func (player *BasicPlayer) GetIP() string {
	return player.connection.GetIP()
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

func (player *BasicPlayer) SendMessage(text *component.TextComponent) {
	player.SendPacket(CreateServerChatMessage(text, 0))
}

func (player *BasicPlayer) Disconnect(text *component.TextComponent) {
	player.SendPacket(CreateServerDisconnect(text))

	player.connection.Remove()
}

func (player *BasicPlayer) SendPacket(packet ServerPacket) {
	player.connection.SendPacket(packet)
}

func (player *BasicPlayer) IsAuthenticated() bool {
	return player.authenticated
}

func (player *BasicPlayer) SetAuthenticated(value bool) {
	player.authenticated = value
}
