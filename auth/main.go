package main

import "auth/server"

func init() {
	server.InitServer()

}

func main() {
	server.StartServer()
}
