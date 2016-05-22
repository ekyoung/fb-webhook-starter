package main

import (
	"bitbucket.org/ekyoung/movebot-fb/server"
)

func main() {
	server := &server.Server{}

	server.Run()
}
