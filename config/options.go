package config

import ()

type Options struct {
	Command    string
	Group      string
	Filter     string
	Stream     string
	StartTime  int64
	EndTime    int64
	Number     int64
	Chunk      int64
	Tail       bool
	TimeZone   bool
	Spacing    bool
	NoGroup    bool
	NoStream   bool
	NoTime     bool
	NoColor    bool
	NoWrap     bool
	TrimLeft   int
	DateFormat string
	Green      string
	Yellow     string
	Red        string
	Refresh    int
	Cache      bool
	CacheDir   string
	Debug      bool
}
