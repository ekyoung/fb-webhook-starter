package main

import (
	"github.com/ekyoung/fb-webhook-starter/server"
)

func main() {
	server := &server.Server{}

	server.Run()
}
