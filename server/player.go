package server

type IPlayer interface {
	GetUsername() string
	SetUsername(string)
	GetUuid() UUID
	SetUuid(UUID)
	GetConnection() IConnection
	SetConnection(IConnection)

	Disconnect(string)
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

func (player *BasicPlayer) Disconnect(reason string) {
	player.SendPacket(&DisconnectPacket{reason: reason})

	player.connection.Remove()
}

func (player *BasicPlayer) SendPacket(packet ServerPacket) {
	player.connection.SendPacket(packet)
}
