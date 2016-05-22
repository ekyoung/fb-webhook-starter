package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct{}

func (s *Server) Run() {
	r := gin.Default()

	r.GET("/webhook", func(c *gin.Context) {
		if c.Query("hub.mode") != "subscribe" {
			c.String(http.StatusOK, "I don't know what you're trying to do.")
			return
		}

		if c.Query("hub.verify_token") != os.ExpandEnv("$VERIFY_TOKEN") {
			c.String(http.StatusOK, "Verify token is incorrect.")
			return
		}

		c.String(http.StatusOK, c.Query("hub.challenge"))
	})

	r.Run()
}
