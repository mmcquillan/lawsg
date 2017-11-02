package config

import (
	"fmt"
	"os"
)

func Env() {

	// parse AWS creds
	if os.Getenv("AWS_REGION") == "" {
		fmt.Println("ERROR: Missing environment variable AWS_REGION")
		os.Exit(1)
	}
	if os.Getenv("AWS_ACCESS_KEY_ID") == "" {
		fmt.Println("ERROR: Missing environment variable AWS_ACCESS_KEY_ID")
		os.Exit(1)
	}
	if os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
		fmt.Println("ERROR: Missing environment variable AWS_AWS_SECRET_ACCESS_KEY")
		os.Exit(1)
	}

}
