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
	if !util.Member(options.Command, "get,groups,streams,help") {
		options.Command = "help"
		options.Group = ""
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

	// validate number and tail
	if options.Tail && options.Number > 0 {
		fmt.Println("ERROR: Cannot tail and limit the number of Log Events")
		os.Exit(1)
	}

	// if
	if options.Cache && options.CacheDir == "" {
		fmt.Println("ERROR: Cannot set Cache without a CacheDir")
		os.Exit(1)
	}

	// don't stat while tailing
	if options.Stats && options.Tail {
		fmt.Println("ERROR: Can't Stat while Tailing logs")
		os.Exit(1)
	}

	// debug out the options
	if options.Debug {
		fmt.Printf("\n==> INPUT: %#v\n", *options)
	}

}
