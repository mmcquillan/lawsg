package config

import (
	"os"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/mmcquillan/lawsg/util"
)

// Defaults - returns the options with defaults set
func Defaults(options *Options) {

	// filter options
	options.Filter = ""
	options.Stream = ""
	options.StartTime = util.ParseDate("10m ago")
	options.EndTime = util.ParseDate("now")
	options.Number = 0
	options.Tail = false

	// display options
	options.TimeZone = false
	options.Spacing = false
	options.NoGroup = false
	options.NoStream = false
	options.NoTime = false
	options.NoColor = false
	options.NoWrap = false
	options.StreamLTrim = 0
	options.StreamRTrim = 0
	options.MessageLTrim = 0
	options.MessageRTrim = 0
	options.DateFormat = ""
	options.Green = ""
	options.Yellow = ""
	options.Red = ""

	// advanced options
	options.Command = ""
	options.Group = ""
	options.Chunk = 10000
	options.Refresh = 5
	options.Cache = false
	home, err := homedir.Dir()
	if err != nil {
		home = os.TempDir()
	}
	options.CacheDir = util.MakePath(home, ".lawsg/cache")
	options.Stats = false
	options.Debug = false

	// stats
	options.Timer = time.Now().UnixNano()

}
