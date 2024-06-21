package server

/// REGION: MinecraftServer

import (
	"time"
)

type MinecraftServer struct {
	listener    IListener
	connFactory IConnectionFactory
	connPool    IConnectionPool
	serverLoop  IServerLoop

	bufferCreate func() IBuffer

	running       bool
	online        bool
	multithreaded bool
}

func CreateMinecraftServer() *MinecraftServer {
	return &MinecraftServer{
		listener:    createBasicListener(),
		connFactory: createBasicConnectionFactory(),
		connPool:    createBasicConnectionPool(),
		serverLoop:  createBasicServerLoop(),
		running:     true,
	}
}

func (server *MinecraftServer) SetListener(listener IListener) {
	server.listener = listener
}

func (server *MinecraftServer) SetConnectionFactory(connFactory IConnectionFactory) {
	server.connFactory = connFactory
}

func (server *MinecraftServer) SetConnectionPool(connPool IConnectionPool) {
	server.connPool = connPool
}

func (server *MinecraftServer) Init() {
	/// please set your custom factories before calling init!!!
	Info("Loading server... (v" + VERSION + ")")
}

func (server *MinecraftServer) SetMojangAuth(value bool) {
	server.online = value
}

func (server *MinecraftServer) SetMultiThreading(value bool) {
	server.multithreaded = value
}

func (server *MinecraftServer) SetBufferCreator(f func() IBuffer) {
	server.bufferCreate = f
}

func (server *MinecraftServer) CreateBuffer() IBuffer {
	return server.bufferCreate()
}

func (server *MinecraftServer) Start(port int) {
	Info("Starting server...")
	server.listener.Start(port, *server)

	Info("Starting logic loop!")

	var lastCall = time.Now().UnixMilli()

	for server.running {
		timez := time.Now().UnixMilli() - lastCall
		lastCall = time.Now().UnixMilli()
		server.serverLoop.Call(timez)
	}
}

func (server *MinecraftServer) Stop() {
	server.running = false
}

/// END;
