package server

/// REGION: MinecraftServer

import (
	"net/deltamc/server/crypto"
	"net/deltamc/server/thread"
	"time"
)

type MinecraftServer struct {
	listener      IListener
	connFactory   IConnectionFactory
	connPool      IConnectionPool
	serverLoop    IServerLoop
	packetHandler IPacketHandler
	authenticator IAuthenticator
	mapper        *ProtocolTable
	serverThread  thread.Thread

	bufferCreate   func() IBuffer
	playerCreate   func(string) IPlayer
	responseCreate func(*MinecraftServer) *ServerResponse

	verificationToken func() []byte
	keypair           *crypto.Keypair

	running       bool
	multithreaded bool
	ticks         int

	favicon string

	injectionManager *InjectionManager
	materialRegistry *MaterialRegistry
}

func CreateMinecraftServer() *MinecraftServer {
	return &MinecraftServer{
		listener:          createBasicListener(),
		connFactory:       createBasicConnectionFactory(),
		connPool:          createBasicConnectionPool(),
		serverLoop:        createBasicServerLoop(),
		packetHandler:     CreatePacketHandler(),
		bufferCreate:      createBasicBuffer,
		playerCreate:      createBasicPlayer,
		responseCreate:    CreateServerResponse,
		verificationToken: GenerateVerificationToken,
		authenticator:     nil,
		mapper:            CreateProtocolTable(),
		keypair:           crypto.NewKeypair(),
		injectionManager:  CreateInjectionManager(),
		materialRegistry:  materials,
		running:           true,
		multithreaded:     false,
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

	server.materialRegistry.Load("")
	server.favicon = loadIcon()
}

func (server *MinecraftServer) SetMultiThreading(value bool) {
	server.multithreaded = value
}

func (server *MinecraftServer) SetProtocolHandler(handler IPacketHandler) {
	server.packetHandler = handler
}

func (server *MinecraftServer) SetKeypair(keypair *crypto.Keypair) {
	server.keypair = keypair
}

func (server *MinecraftServer) GetKeypair() *crypto.Keypair {
	return server.keypair
}

func (server *MinecraftServer) GetFavicon() string {
	return server.favicon
}

func (server *MinecraftServer) NewVerificationToken() []byte {
	return server.verificationToken()
}

func (server *MinecraftServer) SetVerificationTokenFactory(f func() []byte) {
	server.verificationToken = f
}

func (server *MinecraftServer) SetBufferCreator(f func() IBuffer) {
	server.bufferCreate = f
}

func (server *MinecraftServer) SetPlayerCreator(f func(string) IPlayer) {
	server.playerCreate = f
}

func (server *MinecraftServer) GetAuthenticator() IAuthenticator {
	return server.authenticator
}

func (server *MinecraftServer) SetAuthenticator(authenticator IAuthenticator) {
	server.authenticator = authenticator
}

func (server *MinecraftServer) GetInjectionManager() *InjectionManager {
	return server.injectionManager
}

func (server *MinecraftServer) CreateBuffer() IBuffer {
	return server.bufferCreate()
}

func (server *MinecraftServer) Start(port int) {
	Info("Starting server... %d", getTime())
	server.listener.Start(port, server)

	test(server)

	Info("Starting logic loop!")

	var lastCall = time.Now().UnixMilli()

	//test(server)
	server.serverThread = thread.New()

	server.serverThread.CallNonBlock(func() {
		for server.running {
			timez := time.Now().UnixMilli() - lastCall
			lastCall = time.Now().UnixMilli()
			server.serverLoop.Call(timez, server)
		}
	})

	for {
		time.Sleep(time.Millisecond * 10)
	}
}

func (server *MinecraftServer) Stop() {
	server.running = false
}

/// END;
