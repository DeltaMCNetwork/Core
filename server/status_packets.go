package server

import (
	"encoding/json"
	"net/deltamc/server/status"
)

type ServerStatusResponse struct {
	ServerPacket
	Response *status.Response
}

func CreateServerStatusResponse(response *status.Response) *ServerStatusResponse {
	return &ServerStatusResponse{Response: response}
}

func (r *ServerStatusResponse) GetPacketId(conn IConnection) int {
	return 0x00
}

func (r *ServerStatusResponse) Write(buf IBuffer) {
	arr, err := json.Marshal(r.Response)

	if err != nil {
		Fatal("Error in serializing response packet %s", err.Error())
	}

	buf.WriteByteArray(arr)
}

type ServerStatusPong struct {
	ServerPacket
	Payload int64
}

func CreateServerStatusPong(payload int64) *ServerStatusPong {
	return &ServerStatusPong{Payload: payload}
}

func (p *ServerStatusPong) GetPacketId(conn IConnection) int {
	return 0x01
}

func (p *ServerStatusPong) Write(buf IBuffer) {
	buf.WriteLong(p.Payload)
}

var _ IPacket = (*ServerStatusResponse)(nil)
var _ IPacket = (*ServerStatusPong)(nil)
