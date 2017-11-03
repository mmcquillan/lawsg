package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/araddon/dateparse"
	"gopkg.in/alecthomas/kingpin.v2"
)

type Options struct {
	Group      string
	Filter     string
	Stream     string
	StartTime  int64
	EndTime    int64
	Number     int64
	Chunk      int64
	Tail       bool
	TimeZone   bool
	NoGroup    bool
	NoStream   bool
	NoTime     bool
	NoColor    bool
	NoWrap     bool
	TrimLeft   int
	Last       int
	DateFormat string
	Green      string
	Yellow     string
	Red        string
	Debug      bool
}

func Parse() (options Options) {

	// set flags
	group := kingpin.Arg("group", "Log Group or 'list' to show groups").Required().String()
	filter := kingpin.Flag("filter", "Cloudwatch Filter Pattern").Short('f').String()
	stream := kingpin.Flag("stream", "Stream Filter (comma seperated)").String()
	startTime := kingpin.Flag("starttime", "Start Time").Short('s').String()
	endTime := kingpin.Flag("endtime", "End Time").Short('e').String()
	last := kingpin.Flag("last", "Last X minutes of logs").Short('l').Int()
	number := kingpin.Flag("number", "Number of Rows").Short('n').Int64()
	chunk := kingpin.Flag("chunk", "Chunk Size").Int64()
	noGroup := kingpin.Flag("ng", "No Group").Bool()
	noStream := kingpin.Flag("ns", "No Streams").Bool()
	noWrap := kingpin.Flag("nw", "No Wrapping Lines (will be truncated)").Bool()
	noColor := kingpin.Flag("nc", "No Color").Bool()
	noTime := kingpin.Flag("nt", "No Time").Bool()
	timeZone := kingpin.Flag("tz", "Convert to my Time Zone").Short('z').Bool()
	tail := kingpin.Flag("tail", "Tail of Log (experimental)").Short('t').Bool()
	trimLeft := kingpin.Flag("trimleft", "Trim Left of Event Message").Int()
	dateFormat := kingpin.Flag("dateformat", "Date Format for the timestamp (https://golang.org/src/time/format.go)").String()
	green := kingpin.Flag("green", "Green Highlight (comma seperated)").String()
	yellow := kingpin.Flag("yellow", "Yellow Highlight (comma seperated)").String()
	red := kingpin.Flag("red", "Red Highlight (comma seperated)").String()
	debug := kingpin.Flag("debug", "Debug").Bool()
	kingpin.Parse()

	// assign values
	options.Group = *group
	options.Filter = setString("", "LAWSG_FILTER", *filter)
	options.Stream = setString("", "LAWSG_STREAM", *stream)
	options.StartTime = setDate(time.Now().Add(-10*time.Minute).Unix()*1000, "LAWSG_STARTTIME", *startTime)
	options.EndTime = setDate(time.Now().Unix()*1000, "LAWSG_ENDTIME", *endTime)
	options.Last = setInt(0, "LAWSG_LAST", *last)
	options.Number = setInt64(0, "LAWSG_NUMBER", *number)
	options.Chunk = setInt64(10000, "LAWSG_CHUNK", *chunk)
	options.Tail = setBool(false, "LAWSG_TAIL", *tail)
	options.TimeZone = setBool(false, "LAWSG_TIMEZONE", *timeZone)
	options.NoTime = setBool(false, "LAWSG_NO_TIME", *noTime)
	options.NoGroup = setBool(false, "LAWSG_NO_GROUP", *noGroup)
	options.NoStream = setBool(false, "LAWSG_NO_STREAM", *noStream)
	options.NoColor = setBool(false, "LAWSG_NO_COLOR", *noColor)
	options.NoWrap = setBool(false, "LAWSG_NOWRAP", *noWrap)
	options.TrimLeft = setInt(0, "LAWSG_TRIM_LEFT", *trimLeft)
	options.DateFormat = setString("", "LAWSG_DATE_FORMAT", *dateFormat)
	options.Green = setString("", "LAWSG_GREEN", *green)
	options.Yellow = setString("", "LAWSG_YELLOW", *yellow)
	options.Red = setString("", "LAWSG_RED", *red)
	options.Debug = setBool(false, "LAWSG_DEBUG", *debug)

	// run some validation
	if options.EndTime < options.StartTime {
		fmt.Println("ERROR: Start Time is before End Time")
		os.Exit(1)
	}
	if options.Last > 0 {
		options.StartTime = time.Now().Unix()*1000 - int64(options.Last*60*1000)
		options.EndTime = time.Now().Unix() * 1000
	}

	return options
}

func setString(def string, env string, flg string) (val string) {
	val = def
	if os.Getenv(env) != "" {
		val = os.Getenv(env)
	}
	if flg != "" {
		val = flg
	}
	return val
}

func setDate(def int64, env string, flg string) (val int64) {
	val = def
	if os.Getenv(env) != "" {
		if t, err := dateparse.ParseIn(os.Getenv(env), time.UTC); err == nil {
			val = t.Unix() * 1000
		}
	}
	if flg != "" {
		if t, err := dateparse.ParseIn(flg, time.UTC); err == nil {
			val = t.Unix() * 1000
		}
	}
	return val
}

func setInt(def int, env string, flg int) (val int) {
	val = def
	if os.Getenv(env) != "" {
		if v, err := strconv.ParseInt(os.Getenv(env), 10, 32); err == nil {
			val = int(v)
		}
	}
	if flg != 0 {
		val = flg
	}
	return val
}

func setInt64(def int64, env string, flg int64) (val int64) {
	val = def
	if os.Getenv(env) != "" {
		if v, err := strconv.ParseInt(os.Getenv(env), 10, 64); err == nil {
			val = v
		}
	}
	if flg != 0 {
		val = flg
	}
	return val
}

func setBool(def bool, env string, flg bool) (val bool) {
	val = def
	if os.Getenv(env) != "" {
		if v, err := strconv.ParseBool(os.Getenv(env)); err == nil {
			val = v
		}
	}
	if flg != false {
		val = flg
	}
	return val
}
