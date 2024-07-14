package server

type IServerLoop interface {
	Call(time int64, server MinecraftServer)
}

type BasicServerLoop struct {
	IServerLoop
}

func (loop *BasicServerLoop) Call(timeBetween int64, server MinecraftServer) {
	Info("Time between ticks is %dms", timeBetween)

	for _, element := range server.connPool.GetConnections() {
		element.Read(server)
	}
}

func createBasicServerLoop() IServerLoop {
	return &BasicServerLoop{}
}
