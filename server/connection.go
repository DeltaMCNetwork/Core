package server

import (
	"crypto/aes"
	"crypto/cipher"
	"net"
	"net/deltamc/server/crypto"
)

type IConnectionFactory interface {
	CreateConnection(*net.TCPConn, *MinecraftServer)
}

type IConnection interface {
	Read(*MinecraftServer) error
	ReadPacket([]byte, int, *MinecraftServer)
	GetPacketMode() byte
	SetPacketMode(byte)
	GetConnection() *net.TCPConn
	GetCompressionThreshold() int
	SetCompressionThreshold(int)
	GetPlayer() IPlayer
	SetPlayer(IPlayer)
	SendPacket(ServerPacket)
	GetMinecraftServer() *MinecraftServer
	SetMinecraftServer(*MinecraftServer)
	GetProtocolVersion() int
	SetProtocolVersion(int)
	EnableEncryption(secret []byte)
	Remove()
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
	server.connPool.AddConnection(&BasicConnection{
		conn:             conn,
		mode:             PacketModeHandshake,
		threshold:        0,
		server:           server,
		encrptionEnabled: false,
	}, server)
}

func createBasicConnectionFactory() IConnectionFactory {
	return &BasicConnectionFactory{}
}

type BasicConnection struct {
	IConnection
	conn      *net.TCPConn
	mode      byte
	threshold int
	player    IPlayer
	server    *MinecraftServer
	protocol  int

	// encryption stuff
	encrptionEnabled bool
	encrypter        cipher.Stream
	decrypter        cipher.Stream
}

func (conn *BasicConnection) SetProtocolVersion(version int) {
	conn.protocol = version
}

func (conn *BasicConnection) GetProtocolVersion() int {
	return conn.protocol
}

func (conn *BasicConnection) SetPlayer(player IPlayer) {
	conn.player = player
}

func (conn *BasicConnection) GetPlayer() IPlayer {
	return conn.player
}

// Read implements IConnection.
func (conn *BasicConnection) Read(server *MinecraftServer) error {
	data := make([]byte, BUFFER_SIZE)

	size, err := conn.conn.Read(data)

	// get rid of the unused space
	data = data[:size]

	if err != nil {
		conn.Remove()
		Info("Removed connection, there are now %d connections", server.connPool.GetConnectionCount())

		return err
	}

	if size == 0 {
		Debug("Packet is length 0")
	}

	if conn.encrptionEnabled {
		// encryption enabled

		conn.decrypt(data)
	}

	index := 0
	packetsRead := 1

	for index < size {
		// read var int from buffer to get packet length
		packetLength, len := ReadVarInt(data[index:])

		if int(packetLength)+index >= BUFFER_SIZE {
			extraLength := (int(packetLength) + index) - BUFFER_SIZE
			newData := make([]byte, extraLength)
			rest, err := conn.conn.Read(newData)

			if err != nil {
				conn.Remove()

				return err
			}

			if rest != extraLength {
				Error("Expected %d bytes, only received %d bytes.", extraLength, rest)
				conn.Remove()
			}

			if conn.encrptionEnabled {
				// encryption enabled

				conn.decrypt(data)
			}

			data = append(data, newData...)
		}

		//Debug("Reading packet #%d of sucession", packetsRead)

		packetsRead++

		if conn.GetPacketMode() == PacketModePlay && packetLength-int32(len) == 0 {
			index += int(packetLength)
			continue
		}

		packetData := data[index+len : index+len+int(packetLength)]
		index += int(packetLength) + len

		conn.ReadPacket(packetData, int(packetLength), server)
	}

	return nil
}

func (conn *BasicConnection) SendPacket(packet ServerPacket) {
	buf := conn.server.CreateBuffer()
	packet.Write(buf)

	data := buf.GetBytes()
	newBuf := conn.server.CreateBuffer()

	length := len(data) + 1
	Info("Length of packet is %d, and ID is %d", length, packet.GetPacketId(conn))

	newBuf.WriteVarInt(int32(length))
	newBuf.WriteVarInt(int32(packet.GetPacketId(conn)))
	newBuf.Write(data)

	if conn.encrptionEnabled {
		conn.encrypt(newBuf.GetBytes())
	}

	_, err := conn.conn.Write(newBuf.GetBytes())

	if err != nil {
		conn.Remove()
		Info("Client disconnected: %s", err.Error())
	}
}

func (conn *BasicConnection) Remove() {
	conn.conn.Close()
	conn.server.connPool.RemoveConnection(conn)
}

func (conn *BasicConnection) GetMinecraftServer() *MinecraftServer {
	return conn.server
}

func (conn *BasicConnection) SetMinecraftServer(server *MinecraftServer) {
	conn.server = server
}

func (conn *BasicConnection) GetConnection() *net.TCPConn {
	return conn.conn
}

func (conn *BasicConnection) SetCompressionThreshold(value int) {
	conn.threshold = value
}

func (conn *BasicConnection) GetCompressionThreshold() int {
	return conn.threshold
}

func (conn *BasicConnection) ReadPacket(data []byte, length int, server *MinecraftServer) {
	//Info("Packet length is %d", length)

	buf := server.CreateBuffer()
	buf.SetData(data)

	if USE_COMPRESSION && conn.mode == PacketModePlay {
		Info("Compression enabled")
	}

	packetId := buf.ReadVarInt()
	//Info("PacketId is %d and length is %d", packetId, length)

	server.mapper.HandlePacket(packetId, buf, conn, server)
}

func (conn *BasicConnection) GetPacketMode() byte {
	return conn.mode
}

func (conn *BasicConnection) SetPacketMode(value byte) {
	conn.mode = value
}

func (conn *BasicConnection) EnableEncryption(secret []byte) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		Error("Error creating cipher: %s", err.Error())
		conn.Remove()
		return
	}

	conn.decrypter = crypto.NewCFB8Decrypter(block, secret)
	conn.encrypter = crypto.NewCFB8Encrypter(block, secret)

	conn.encrptionEnabled = true
}

func (conn *BasicConnection) decrypt(bytearr []byte) {
	conn.decrypter.XORKeyStream(bytearr, bytearr)
}

func (conn *BasicConnection) encrypt(bytearr []byte) {
	conn.encrypter.XORKeyStream(bytearr, bytearr)
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
