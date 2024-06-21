package main

import "org/example/server"

func main() {
	//this is a tests file

	serv := server.CreateMinecraftServer()

	serv.SetListener(&server.BasicListener{})

	serv.Init()
	serv.Start(25565)
}
