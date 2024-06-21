package server

import "net"

type IConnectionFactory interface {
	CreateConnection(net.TCPConn, MinecraftServer)
}

type IConnection interface {
}

type IConnectionPool interface {
	Tick(MinecraftServer)
	GetConnections() []IConnection
	GetConnectionCount() int
	AddConnection(IConnection, MinecraftServer)
}

type BasicConnectionFactory struct {
	IConnectionFactory
}

func (factory *BasicConnectionFactory) CreateConnection(conn net.TCPConn, server MinecraftServer) {
	server.connPool.AddConnection(BasicConnection{
		conn: conn,
	}, server)
}

func createBasicConnectionFactory() IConnectionFactory {
	return &BasicConnectionFactory{}
}

type BasicConnection struct {
	IConnection
	conn net.TCPConn
}

type BasicConnectionPool struct {
	IConnectionPool
	connections []IConnection
}

func (pool *BasicConnectionPool) Tick(server MinecraftServer) {

}

func (pool *BasicConnectionPool) GetConnections() []IConnection {
	return pool.connections
}

func (pool *BasicConnectionPool) GetConnectionCount() int {
	return len(pool.connections)
}

func (pool *BasicConnectionPool) AddConnection(conn IConnection, server MinecraftServer) {
	pool.connections = append(pool.connections, conn)
}

func createBasicConnectionPool() IConnectionPool {
	return &BasicConnectionPool{
		connections: make([]IConnection, 0),
	}
}
