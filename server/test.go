package server

import "os"

func test(server *MinecraftServer) {
	//nbt test

	/*nbtData*/
	nbtData, err := os.ReadFile("resources/scoreboard.dat")

	if err != nil {
		return
	}

	buf := server.CreateBuffer()
	compound := NbtReadGzip(nbtData, buf)

	Info(compound)

	writeBuf := server.CreateBuffer()
	NbtWrite(writeBuf, *compound)

	//Info("Len of writebuf is %d", writeBuf.GetLength())

	newCompound := NbtRead(writeBuf)

	Info(newCompound)

	data := NbtWriteGzip(server.CreateBuffer(), *newCompound)

	os.WriteFile("resources/scoreboard.dat", data, 0666)
}
