package main

import "user/server"

func init() {
	server.InitServer()

}

func main() {
	server.StartServer()
}
