package server

import "net"

type IConnectionFactory interface {
	CreateConnection(*net.TCPConn, *MinecraftServer)
}

type IConnection interface {
	Read(*MinecraftServer) error
	ReadPacket([]byte, int, *MinecraftServer)
	GetPacketMode() int
	SetPacketMode(*int)
	GetConnection() *net.TCPConn
}

type IConnectionPool interface {
	Tick(*MinecraftServer)
	GetConnections() []IConnection
	GetConnectionCount() int
	AddConnection(IConnection, *MinecraftServer)
	RemoveConnection(IConnection)
}

type BasicConnectionFactory struct {
	IConnectionFactory
}

func (factory *BasicConnectionFactory) CreateConnection(conn *net.TCPConn, server *MinecraftServer) {
	server.connPool.AddConnection(BasicConnection{
		conn: conn,
	}, server)
}

func createBasicConnectionFactory() IConnectionFactory {
	return &BasicConnectionFactory{}
}

type BasicConnection struct {
	IConnection
	conn *net.TCPConn
	mode *int
}

// Read implements IConnection.
func (conn BasicConnection) Read(server *MinecraftServer) error {
	data := make([]byte, BUFFER_SIZE)

	size, err := conn.conn.Read(data)

	if err != nil {
		server.connPool.RemoveConnection(conn)

		Info("Removed connection, there are now %d connections", server.connPool.GetConnectionCount())

		return err
	}

	if size == 0 {
		Info("Packet is length 0")
	}

	conn.ReadPacket(data, size, server)

	return nil
}

func (conn BasicConnection) GetConnection() *net.TCPConn {
	return conn.conn
}

func (conn BasicConnection) ReadPacket(data []byte, length int, server *MinecraftServer) {
	Info("Packet length is %d", length)
}

func (conn BasicConnection) GetPacketMode() int {
	return *conn.mode
}

func (conn BasicConnection) SetPacketMode(value *int) {
	conn.mode = value
}

type BasicConnectionPool struct {
	IConnectionPool
	connections []IConnection
}

func (pool *BasicConnectionPool) Tick(server *MinecraftServer) {

}

func (pool *BasicConnectionPool) GetConnections() []IConnection {
	return pool.connections
}

func (pool *BasicConnectionPool) GetConnectionCount() int {
	return len(pool.connections)
}

func (pool *BasicConnectionPool) RemoveConnection(conn IConnection) {
	// find index
	index := -1
	for i := 0; i < pool.GetConnectionCount(); i++ {
		if pool.connections[i] == conn {
			index = i
		}
	}

	if index == -1 {
		return
	}

	pool.connections = append(pool.connections[:index], pool.connections[index+1:]...)
}

func (pool *BasicConnectionPool) AddConnection(conn IConnection, server *MinecraftServer) {
	pool.connections = append(pool.connections, conn)
}

func createBasicConnectionPool() IConnectionPool {
	return &BasicConnectionPool{
		connections: make([]IConnection, 0),
	}
}
