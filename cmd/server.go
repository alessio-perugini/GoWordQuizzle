package main

import "github.com/alessio-perugini/GoWordQuizzle/server"

func main() {
	server.StartRPC()
	server.StartTCP()
}
