package config

import (
	"os"
	"strings"
)

// Commands - parse for commands and groups
func Commands(options *Options) {

	if len(os.Args) > 1 && !strings.HasPrefix(os.Args[1], "-") {
		if options.Command == "" {
			options.Command = os.Args[1]
		}
	}

	if len(os.Args) > 2 && !strings.HasPrefix(os.Args[2], "-") {
		if options.Group == "" {
			options.Group = os.Args[2]
		}
	}

}
