package main

import (
	"log"

	"github.com/ekyoung/configsrc"

	"github.com/ekyoung/fb-webhook-starter/server"
)

type AppConfig struct {
	FacebookWebhookVerifyToken string `yaml:"facebook_webhook_verify_token"`
}

func main() {
	cfg := AppConfig{}

	err := configsrc.Load("movebot_fb_bot", &cfg)
	if err != nil {
		log.Fatalf("error: %v\n", err)
		return
	}

	log.Printf("[DEBUG] Config: %v\n", cfg)

	server := &server.Server{
		FacebookWebhookVerifyToken: cfg.FacebookWebhookVerifyToken,
	}

	server.Run()
}
