package server

func test(server *MinecraftServer) {
	//nbt test

	/*nbtData*/
	nbt := NbtCompound{
		"serverAliveTicks": 9499999,
		"whatTheSigma": NbtCompound{
			"anotherKey": NbtIntArray{
				1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
			},
			"username": "freshday",
		},
	}

	NbtWrite(server.CreateBuffer(), nbt)
}
	