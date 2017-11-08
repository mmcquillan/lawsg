package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mmcquillan/lawsg/util"
	flag "github.com/ogier/pflag"
)

// Flags - pull all of the flagged cli's
func Flags(options *Options) {

	f := flag.NewFlagSet("lawsg", flag.ContinueOnError)

	// filter options
	f.StringVarP(&options.Filter, "filter", "f", options.Filter, "")
	f.StringVarP(&options.Stream, "stream", "m", options.Stream, "")
	var st string
	f.StringVarP(&st, "starttime", "s", "", "")
	var et string
	f.StringVarP(&et, "endtime", "e", "", "")
	f.Int64VarP(&options.Number, "number", "n", options.Number, "")
	f.BoolVarP(&options.Tail, "tail", "t", options.Tail, "")
	var watch bool
	f.BoolVar(&watch, "watch", false, "")

	// display options
	f.BoolVar(&options.TimeZone, "tz", options.TimeZone, "")
	f.BoolVar(&options.Spacing, "spacing", options.Spacing, "")
	f.BoolVar(&options.NoGroup, "ng", options.NoGroup, "")
	f.BoolVar(&options.NoStream, "ns", options.NoStream, "")
	f.BoolVar(&options.NoTime, "nt", options.NoTime, "")
	f.BoolVar(&options.NoColor, "nc", options.NoColor, "")
	f.BoolVar(&options.NoWrap, "nw", options.NoWrap, "")
	f.IntVar(&options.StreamLTrim, "stream-ltrim", options.StreamLTrim, "")
	f.IntVar(&options.StreamRTrim, "stream-rtrim", options.StreamRTrim, "")
	f.IntVar(&options.MessageLTrim, "message-ltrim", options.MessageLTrim, "")
	f.IntVar(&options.MessageRTrim, "message-rtrim", options.MessageRTrim, "")
	f.StringVar(&options.DateFormat, "dateformat", options.DateFormat, "")
	f.StringVar(&options.Green, "green", options.Green, "")
	f.StringVar(&options.Yellow, "yellow", options.Yellow, "")
	f.StringVar(&options.Red, "red", options.Red, "")

	// advanced options
	f.StringVarP(&options.Command, "command", "c", options.Command, "")
	f.StringVarP(&options.Group, "group", "g", options.Group, "")
	f.Int64Var(&options.Chunk, "chunk", options.Chunk, "")
	f.IntVar(&options.Refresh, "refresh", options.Refresh, "")
	f.BoolVar(&options.Cache, "cache", options.Cache, "")
	f.StringVar(&options.CacheDir, "cachedir", options.CacheDir, "")
	f.BoolVar(&options.Stats, "stats", options.Stats, "")
	f.BoolVar(&options.Debug, "debug", options.Debug, "")
	f.SetOutput(ioutil.Discard)
	err := f.Parse(os.Args)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}

	if st != "" {
		options.StartTime = util.ParseDate(st)
	}
	if et != "" {
		options.EndTime = util.ParseDate(et)
	}

	// for cross compatability with other tools
	if watch {
		options.Tail = true
	}

}
