package main

import (
	"simple_api/server"
)

func main() {
	sv := server.Register()

	sv.Start()
}
