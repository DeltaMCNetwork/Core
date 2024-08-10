package server

import (
	"fmt"
	"net"
)

type IListener interface {
	Start(int, *MinecraftServer)
	Stop()
}

type BasicListener struct {
	IListener
}

// Start implements IListener.

// Stop implements IListener.

func (listener *BasicListener) Start(port int, server *MinecraftServer) {
	address, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", "0.0.0.0", port))

	if err != nil {
		return
	}

	socket, err := net.ListenTCP("tcp", address)
	if err != nil {
		return
	}

	go func() {
		for server.running {
			connection, err := socket.AcceptTCP()

			Info("Connection received from %s", connection.RemoteAddr().String())

			if err != nil {
				continue
			}

			conn := server.connFactory.CreateConnection(connection, server)
			// configuration

			conn.Read(server)
		}
	}()
}

func (listener *BasicListener) Stop() {

}

func createBasicListener() IListener {
	return &BasicListener{}
}
