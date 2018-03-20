package config

import ()

type Options struct {

	// filter options
	Filter    string
	Stream    string
	StartTime int64
	EndTime   int64
	Number    int64
	Chunk     int64
	Tail      bool

	// display options
	TimeZone     bool   `json:"time_zone"`
	Spacing      bool   `json:"spacing"`
	NoGroup      bool   `json:"no_group"`
	NoStream     bool   `json:"no_stream"`
	NoTime       bool   `json:"no_time"`
	NoColor      bool   `json:"no_color"`
	NoWrap       bool   `json:"no_wrap"`
	StreamLTrim  int    `json:"stream_ltrim"`
	StreamRTrim  int    `json:"stream_rtrim"`
	MessageLTrim int    `json:"message_ltrim"`
	MessageRTrim int    `json:"message_rtrim"`
	MultiLine    bool   `json:"multi_line"`
	DateFormat   string `json:"date_format"`
	Green        string `json:"green"`
	Yellow       string `json:"yellow"`
	Red          string `json:"red"`

	// advanced options
	Command string
	Group   string
	Refresh int
	Stats   bool
	Debug   bool
	Env     string

	// stats
	Timer int64
}

type SavedOptions map[string]Options
