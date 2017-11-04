package config

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/araddon/dateparse"
	flag "github.com/ogier/pflag"
)

// Flags - pull all of the flagged cli's
func Flags(options *Options) {

	f := flag.NewFlagSet("lawsg", flag.ContinueOnError)

	f.StringVarP(&options.Command, "command", "c", options.Command, "")
	f.StringVarP(&options.Group, "group", "g", options.Group, "")
	f.StringVarP(&options.Filter, "filter", "f", options.Filter, "")
	f.StringVarP(&options.Stream, "stream", "m", options.Stream, "")
	var st string
	flag.StringVarP(&st, "starttime", "s", "", "")
	var et string
	f.StringVarP(&et, "endtime", "e", "", "")
	f.Int64VarP(&options.Number, "number", "n", options.Number, "")
	f.Int64Var(&options.Chunk, "chunk", options.Chunk, "")
	f.BoolVarP(&options.Tail, "tail", "t", options.Tail, "")
	f.BoolVar(&options.TimeZone, "tz", options.TimeZone, "")
	f.BoolVar(&options.Spacing, "spacing", options.Spacing, "")
	f.BoolVar(&options.NoGroup, "ng", options.NoGroup, "")
	f.BoolVar(&options.NoStream, "ns", options.NoStream, "")
	f.BoolVar(&options.NoTime, "nt", options.NoTime, "")
	f.BoolVar(&options.NoColor, "nc", options.NoColor, "")
	f.BoolVar(&options.NoWrap, "nw", options.NoWrap, "")
	f.IntVar(&options.TrimLeft, "trimleft", options.TrimLeft, "")
	f.IntVarP(&options.Last, "last", "l", options.Last, "")
	f.StringVar(&options.DateFormat, "dateformat", options.DateFormat, "")
	f.StringVar(&options.Green, "green", options.Green, "")
	f.StringVar(&options.Yellow, "yellow", options.Yellow, "")
	f.StringVar(&options.Red, "red", options.Red, "")
	f.BoolVar(&options.Debug, "debug", options.Debug, "")
	f.SetOutput(ioutil.Discard)
	f.Parse(os.Args)

	if stp, err := dateparse.ParseIn(st, time.UTC); err == nil {
		options.StartTime = stp.Unix() * 1000
	}

	if etp, err := dateparse.ParseIn(et, time.UTC); err == nil {
		options.EndTime = etp.Unix() * 1000
	}

}
