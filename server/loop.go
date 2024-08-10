package server

type IServerLoop interface {
	Tick(time int64, server *MinecraftServer)
}

type BasicServerLoop struct {
	IServerLoop
	timer *Timer
}

func (loop *BasicServerLoop) Tick(time int64, server *MinecraftServer) {
	if !loop.timer.HasTimePassed(TICK_TIME) {
		return
	}

	server.ticks++

	server.injectionManager.Post(&ServerTickEvent{
		Count: server.ticks,
	})

	loop.timer.Reset()
}

func createBasicServerLoop() IServerLoop {
	return &BasicServerLoop{
		timer: CreateTimer(),
	}
}
