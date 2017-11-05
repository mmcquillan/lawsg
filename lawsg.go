package main

import (
	"github.com/mmcquillan/lawsg/config"
	"github.com/mmcquillan/lawsg/fetch"
)

func main() {
	config.Aws()
	var options config.Options
	config.Defaults(&options)
	config.EnvVars(&options)
	config.Flags(&options)
	config.Commands(&options)
	config.Validate(&options)
	switch options.Command {
	case "groups":
		fetch.Groups(options)
	case "streams":
		fetch.Streams(options)
	case "get":
		fetch.Logs(options)
	default:
		fetch.Help(options)
	}
}
