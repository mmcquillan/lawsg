package main

import (
	"github.com/mmcquillan/lawsg/config"
	"github.com/mmcquillan/lawsg/fetch"
)

func main() {
	// http://docs.aws.amazon.com/sdk-for-go/api/service/cloudwatchlogs
	config.Env()
	options := config.Parse()
	if options.Group == "list" {
		fetch.List()
	} else {
		fetch.Logs(options)
	}
}
