package main

import (
	"fmt"
	"net/deltamc/server"
)

func main() {
	//this is a tests file

	serv := server.CreateMinecraftServer()

	fmt.Println("fuck off")

	serv.SetListener(&server.BasicListener{})
	serv.SetMojangAuth(true)
	serv.SetMultiThreading(true)

	fmt.Println("fuck off")
	serv.Init()
	serv.Start(25565)
}
