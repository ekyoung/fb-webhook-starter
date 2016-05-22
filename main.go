package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"gopkg.in/yaml.v2"

	"github.com/ekyoung/fb-webhook-starter/server"
)

type Config struct {
	FacebookWebhookVerifyToken string `yaml:"facebook_webhook_verify_token"`
}

func main() {
	svc := s3.New(session.New())

	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("movebot"),
		Key:    aws.String("fb-bot/local/ethan-home-config.yaml"),
	})

	if err != nil {
		log.Fatalf("error downloading config: %v\n\n", err)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading GetObject response body: %v\n\n", err)
		return
	}

	cfg := Config{}

	err = yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		log.Fatalf("error: %v\n", err)
		return
	}

	fmt.Printf("Config: \n%v\n\n", cfg)
	//svc := s3.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

	server := &server.Server{
		FacebookWebhookVerifyToken: cfg.FacebookWebhookVerifyToken,
	}

	server.Run()
}
