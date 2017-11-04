package main

import (
	"github.com/mmcquillan/lawsg/config"
	"github.com/mmcquillan/lawsg/fetch"
)

func main() {
	// http://docs.aws.amazon.com/sdk-for-go/api/service/cloudwatchlogs
	config.Aws()
	var options config.Options
	config.Defaults(&options)
	config.EnvVars(&options)
	config.Flags(&options)
	config.Commands(&options)
	config.Validate(&options)
	switch options.Command {
	case "groups":
		fetch.Groups()
	case "streams":
		fetch.Streams(options)
	case "get":
		fetch.Logs(options)
	default:
		fetch.Help(options)
	}
}
