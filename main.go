package main

import "rentcar/webserver"

func main() {
	server := webserver.New()
	server.Run()
}
