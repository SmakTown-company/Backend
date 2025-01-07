package main

import "basket/server"

func init() {
	server.InitServer()
}

func main() {
	server.StartServer()

}
