package server

import (
	"encoding/json"
	"net/deltamc/server/component"
	"net/deltamc/server/crypto"
)

type ServerEncryptionRequest struct {
	ServerPacket
	PublicKey *crypto.PublicKey
	Token     []byte
}

func CreateServerEncryptionRequest(key *crypto.PublicKey, token []byte) *ServerEncryptionRequest {
	return &ServerEncryptionRequest{
		PublicKey: key,
		Token:     token,
	}
}

func (request *ServerEncryptionRequest) GetPacketId(conn IConnection) int {
	return ServerEncryptionRequestPacket
}

func (request *ServerEncryptionRequest) Write(buf IBuffer) {
	// SERVER ID IS FIRST

	buf.WriteString("")

	buf.WriteVarInt(int32(request.PublicKey.Len))
	buf.Write(request.PublicKey.Key)
	buf.WriteVarInt(VERIFY_TOKEN_LENGTH)
	buf.Write(request.Token)
}

type ServerDisconnect struct {
	ServerPacket
	Reason *component.TextComponent // read the write code
}

func CreateServerDisconnect(text *component.TextComponent) *ServerDisconnect {
	// the bold fields are unused in the client protocol
	// the text components don't use that data
	// the client renders it using the & signs
	// well use the formatting color codes

	return &ServerDisconnect{
		Reason: text,
	}
}

func (packet *ServerDisconnect) GetPacketId(conn IConnection) int {
	if conn.GetPacketMode() == PacketModeLogin {
		return ServerLoginDisconnectPacket
	}

	return ServerDisconnectPacket
}

func (packet *ServerDisconnect) Write(buf IBuffer) {
	data, err := json.Marshal(packet.Reason)

	if err != nil {
		Error("Error serializing disconnect packet! %s", err.Error())

		return
	}

	jsonData := string(data)
	//Debug("JSON data is %s", jsonData)

	buf.WriteString(jsonData)
}

type ServerLoginSuccess struct {
	ServerPacket
	UUID     string
	Username string
}

func CreateServerLoginSuccess(uuid, username string) *ServerLoginSuccess {
	return &ServerLoginSuccess{
		UUID:     uuid,
		Username: username,
	}
}

func (packet *ServerLoginSuccess) GetPacketId(conn IConnection) int {
	return ServerLoginSuccessPacket
}

func (packet *ServerLoginSuccess) Write(buf IBuffer) {
	buf.WriteString(packet.UUID)
	buf.WriteString(packet.Username)
}
