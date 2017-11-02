package config

import (
	"log"
	"os"
)

func Env() {

	// parse AWS creds
	if os.Getenv("AWS_ACCESS_KEY_ID") == "" {
		log.Fatal("Missing environment variable: AWS_ACCESS_KEY_ID")
	}
	if os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
		log.Fatal("Missing environment variable: AWS_AWS_SECRET_ACCESS_KEY")
	}

}
