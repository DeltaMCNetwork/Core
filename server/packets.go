package server

import "encoding/json"

type IPacket interface {
	GetPacketId(IConnection) int
}

type ServerPacket interface {
	Write(IBuffer) IBuffer
	IPacket
}

type ClientPacket interface {
	Read(IBuffer, IConnection, *MinecraftServer) *IPacket
	IPacket
}

type DisconnectPacket struct {
	ServerPacket
	reason string
}

func CreateDisconnectPacket(reason string) *DisconnectPacket {
	return &DisconnectPacket{
		reason: reason,
	}
}

func (packet *DisconnectPacket) GetPacketId(conn IConnection) int {
	if conn.GetPacketMode() == PacketModeLogin {
		return 0
	}

	return 0x40
}

func (packet *DisconnectPacket) Write(buf IBuffer) IBuffer {
	data, err := json.Marshal(Message{Text: packet.reason})

	if err != nil {
		Error("Error serializing disconnect packet! %s", err.Error())

		return buf
	}

	jsonData := string(data)
	//Debug("JSON data is %s", jsonData)

	buf.WriteString(jsonData)

	return buf
}
