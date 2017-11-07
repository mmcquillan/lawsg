package config

import (
	"github.com/mmcquillan/lawsg/util"
)

// Defaults - returns the options with defaults set
func Defaults(options *Options) {
	options.Command = ""
	options.Group = ""
	options.Filter = ""
	options.Stream = ""
	options.StartTime = util.ParseDate("10m ago")
	options.EndTime = util.ParseDate("now")
	options.Number = 0
	options.Chunk = 10000
	options.Tail = false
	options.TimeZone = false
	options.Spacing = false
	options.NoGroup = false
	options.NoStream = false
	options.NoTime = false
	options.NoColor = false
	options.NoWrap = false
	options.TrimLeft = 0
	options.DateFormat = ""
	options.Green = ""
	options.Yellow = ""
	options.Red = ""
	options.Refresh = 5
	options.Cache = false
	options.CacheDir = "/tmp/.lawsg"
	options.Debug = false
}
