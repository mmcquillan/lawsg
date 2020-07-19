package config

import (
	"fmt"
	"os"
)

// Aws - checks that AWS creds exist
func Aws(options *Options) {

	if os.Getenv("AWS_SDK_LOAD_CONFIG") == "" {

		if options.Env != "" {
			os.Setenv("AWS_REGION", os.Getenv(options.Env+"_AWS_REGION"))
			os.Setenv("AWS_ACCESS_KEY_ID", os.Getenv(options.Env+"_AWS_ACCESS_KEY_ID"))
			os.Setenv("AWS_SECRET_ACCESS_KEY", os.Getenv(options.Env+"_AWS_SECRET_ACCESS_KEY"))
		}

		if options.Region != "" {
			os.Setenv("AWS_REGION", options.Region)
		}

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

	} else {

		if options.Env != "" {
			os.Setenv("AWS_DEFAULT_PROFILE", options.Env)
		}

		if options.Region != "" {
			os.Setenv("AWS_REGION", options.Region)
		}

		if os.Getenv("AWS_REGION") == "" {
			fmt.Println("ERROR: Missing environment variable AWS_REGION")
			os.Exit(1)
		}

	}

}
