package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	FacebookWebhookVerifyToken string
}

func (s *Server) Run() {
	r := gin.Default()

	r.GET("/webhook", func(c *gin.Context) {
		if c.Query("hub.mode") != "subscribe" {
			c.String(http.StatusOK, "I don't know what you're trying to do.")
			return
		}

		if c.Query("hub.verify_token") != s.FacebookWebhookVerifyToken {
			c.String(http.StatusOK, "Verify token is incorrect.")
			return
		}

		c.String(http.StatusOK, c.Query("hub.challenge"))
	})

	r.Run()
}
