package server

type IServerLoop interface {
	Call(time int64, server MinecraftServer)
	Tick(time int64, server MinecraftServer)
}

type BasicServerLoop struct {
	IServerLoop
	timer *Timer
}

func (loop *BasicServerLoop) Call(timeBetween int64, server MinecraftServer) {
	//Info("Time between Loop is %dms", timeBetween)

	//loop stuff

	for _, element := range server.connPool.GetConnections() {
		element.Read(server)
	}

	//end loop stuff

	if loop.timer.HasTimePassed(TICK_TIME) {
		loop.Tick(loop.timer.GetPassed(), server)
	}
}

func (loop *BasicServerLoop) Tick(time int64, server MinecraftServer) {
	Info("Time passed between server ticks: %d", time)

	loop.timer.Reset()
}

func createBasicServerLoop() IServerLoop {
	return &BasicServerLoop{
		timer: CreateTimer(),
	}
}
