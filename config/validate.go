package config

import (
	"fmt"
	"os"

	"github.com/mmcquillan/lawsg/util"
)

func Validate(options *Options) {

	// spelling
	if options.Command == "group" {
		fmt.Println("ERROR: Did you mean 'groups'?")
		os.Exit(1)
	}
	if options.Command == "stream" {
		fmt.Println("ERROR: Did you mean 'streams'?")
		os.Exit(1)
	}

	// validate command
	if options.Command != "get" && options.Command != "groups" && options.Command != "streams" {
		options.Command = "help"
	}

	// validate group
	if options.Command == "get" || options.Command == "streams" {
		if options.Group == "" {
			fmt.Println("ERROR: You must specify a group")
			os.Exit(1)
		}
	}

	// validate timing
	if options.EndTime < options.StartTime {
		fmt.Println("ERROR: Start Time is before End Time")
		os.Exit(1)
	}
	if options.Last > 0 {
		options.StartTime = util.ParseDate("10m ago")
		options.EndTime = util.ParseDate("now")
	}

	// debug out the options
	if options.Debug {
		fmt.Printf("\n==> INPUT: %#v\n", *options)
	}

}
