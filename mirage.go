package main

import (
	"server"
)

func main(){
	server := new(server.Server)
	server.Start()
}
