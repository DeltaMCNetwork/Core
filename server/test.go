package server

func test(server *MinecraftServer) {
	server.injectionManager.Register(func(event *ServerTickEvent) {
		Info("Tick Count: %d", event.Count)

		event.Count = 999
	})

	server.injectionManager.Register(func(event *ServerTickEvent) {
		Info("Test %d", event.Count)
	})
}
