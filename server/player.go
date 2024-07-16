package server

type IPlayer interface {
	GetUsername() string
	SetUsername(string)
	GetUuid() UUID
	SetUuid(UUID)
}

type BasicPlayer struct {
	IPlayer
	username string
	uuid     UUID
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
