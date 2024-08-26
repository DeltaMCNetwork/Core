package server

func test(server *MinecraftServer) {
	//nbt test

	//bit shit

	blockId := 1763
	metadata := 7

	var bytes = []uint8{uint8(blockId << 8), uint8((blockId) | metadata)}

	Info("val is %d:%d", bytes[0], bytes[1])
}
