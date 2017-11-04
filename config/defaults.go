package config

import (
	"time"
)

// Defaults - returns the options with defaults set
func Defaults(options *Options) {
	options.Command = ""
	options.Group = ""
	options.Filter = ""
	options.Stream = ""
	options.StartTime = time.Now().Add(-10*time.Minute).Unix() * 1000
	options.EndTime = time.Now().Unix() * 1000
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
	options.Last = 0
	options.DateFormat = ""
	options.Green = ""
	options.Yellow = ""
	options.Red = ""
	options.Debug = false
}
