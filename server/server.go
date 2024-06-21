package server

/// REGION: MinecraftServer

import (
	"time"

	"github.com/charmbracelet/log"
)

type MinecraftServer struct {
	listener    IListener
	connFactory IConnectionFactory
	connPool    IConnectionPool
	serverLoop  IServerLoop

	running bool
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
	log.Info("Loading server... (v" + VERSION + ")")
}

func (server *MinecraftServer) Start(port int) {
	log.Info("Starting server...")
	server.listener.Start(port, *server)

	log.Info("Starting logic loop!")

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
