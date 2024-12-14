package main

import (
	"card/server"
)

func init() {
	server.InitServer()

}

func main() {
	server.StartServer()
}
