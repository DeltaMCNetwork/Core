package main

import (
	"net/deltamc/server"
)

func main() {
	//this is a tests file

	serv := server.CreateMinecraftServer()

	serv.SetListener(&server.BasicListener{})
	serv.SetMultiThreading(true)
	serv.SetAuthenticator(server.CreateMojangAuthenticator())

	serv.Init()
	serv.Start(25565)
}
