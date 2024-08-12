package server

import (
	"math/rand"
	"net/deltamc/server/component"
	"time"
)

type IKeepAliveSender interface {
	StartSendingKeepalive(*MinecraftServer)
	ResetCounter(connection IConnection) // called when a keepalive is received
}

type BasicKeepaliveSender struct {
	lastReceivedKeepalive map[IConnection]time.Time
	lastSentKeepalive     map[IConnection]time.Time
}

func CreateBasicKeepAliveSender() *BasicKeepaliveSender {
	return &BasicKeepaliveSender{
		lastReceivedKeepalive: make(map[IConnection]time.Time),
		lastSentKeepalive:     make(map[IConnection]time.Time),
	}
}

func (b *BasicKeepaliveSender) StartSendingKeepalive(server *MinecraftServer) {
	go func() {
		for server.running {
			for _, connection := range server.connPool.GetConnections() {
				if !(connection.GetPacketMode() == PacketModePlay) {
					continue
				}

				lastSent, ok := b.lastSentKeepalive[connection]
				if !ok {
					// new connection
					b.sendKeepalive(connection)
					continue
				}

				if time.Since(lastSent) > KEEPALIVE_SEND_TIME*time.Second {
					// time to send another keepalive
					b.sendKeepalive(connection)
				}

				lastReceived, ok := b.lastReceivedKeepalive[connection]
				if !ok {
					b.ResetCounter(connection)
					continue
				}

				if time.Since(lastReceived) > KEEPALIVE_TIMEOUT_TIME*time.Second {
					// timeout
				}
			}
		}
	}()
}

func (b *BasicKeepaliveSender) sendKeepalive(connection IConnection) {
	b.lastSentKeepalive[connection] = time.Now()
	packet := CreateServerKeepAlive(int32(rand.Int()))
	connection.SendPacket(packet)
}

func (b *BasicKeepaliveSender) ResetCounter(connection IConnection) {
	b.lastReceivedKeepalive[connection] = time.Now()
}

func (b *BasicKeepaliveSender) timeout(connection IConnection) {
	// remove the connection from the tracker
	delete(b.lastSentKeepalive, connection)
	delete(b.lastReceivedKeepalive, connection)

	// disconnect them
	connection.GetPlayer().Disconnect(component.NewTextComponent("Timed out"))
}

var _ IKeepAliveSender = (*BasicKeepaliveSender)(nil)
