package server

func test(server *MinecraftServer) {
	server.injectionManager.Register(func(event *ServerTickEvent) {
		//Info("Tick Count: %d", event.Count)

		event.Count = 999
	})

	server.injectionManager.Register(func(event *ServerTickEvent) {
		//Info("Test %d", event.Count)
	})

	server.injectionManager.Register(func(event *PacketReceivedEvent) {
		//Info("Received packet, cancelling.")

		event.cancelled = true
	})

	if server.injectionManager.Post(&PacketReceivedEvent{}) {
		Info("Cancelled")
	}
}
