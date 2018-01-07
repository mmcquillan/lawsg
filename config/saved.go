package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/mitchellh/go-homedir"
	"github.com/mmcquillan/lawsg/util"
)

// Saved - loads any saved configs by gets
func Saved(options *Options) {
	if options.Command == "get" {
		home, err := homedir.Dir()
		if err != nil {
			return
		}
		conf := ""
		saveFile := false

		// test .lawsg
		conf = util.MakePath(home, ".lawsg")
		if util.FileExists(conf) {
			saveFile = true
		}

		// test .lawsg.json
		conf = util.MakePath(home, ".lawsg.json")
		if util.FileExists(conf) {
			saveFile = true
		}

		if saveFile {
			file, err := ioutil.ReadFile(conf)
			if err != nil {
				fmt.Println("WARN: Unable to read Config File - ", err)
				return
			}
			var savedOptions SavedOptions
			err = json.Unmarshal(file, &savedOptions)
			if err != nil {
				fmt.Println("WARN: Config File Unmarshal - ", err)
				return
			}
			if o, exists := savedOptions["global"]; exists {
				options.TimeZone = setBool(options.TimeZone, o.TimeZone)
				options.Spacing = setBool(options.Spacing, o.Spacing)
				options.NoGroup = setBool(options.NoGroup, o.NoGroup)
				options.NoStream = setBool(options.NoStream, o.NoStream)
				options.NoTime = setBool(options.NoTime, o.NoTime)
				options.NoColor = setBool(options.NoColor, o.NoColor)
				options.NoWrap = setBool(options.NoWrap, o.NoWrap)
				options.StreamLTrim = setInt(options.StreamLTrim, o.StreamLTrim)
				options.StreamRTrim = setInt(options.StreamRTrim, o.StreamRTrim)
				options.MessageLTrim = setInt(options.MessageLTrim, o.MessageLTrim)
				options.MessageRTrim = setInt(options.MessageRTrim, o.MessageRTrim)
				options.MultiLine = setBool(options.MultiLine, o.MultiLine)
				options.DateFormat = setString(options.DateFormat, o.DateFormat)
				options.Green = setString(options.Green, o.Green)
				options.Yellow = setString(options.Yellow, o.Yellow)
				options.Red = setString(options.Red, o.Red)
			}
			if o, exists := savedOptions[options.Group]; exists {
				options.TimeZone = setBool(options.TimeZone, o.TimeZone)
				options.Spacing = setBool(options.Spacing, o.Spacing)
				options.NoGroup = setBool(options.NoGroup, o.NoGroup)
				options.NoStream = setBool(options.NoStream, o.NoStream)
				options.NoTime = setBool(options.NoTime, o.NoTime)
				options.NoColor = setBool(options.NoColor, o.NoColor)
				options.NoWrap = setBool(options.NoWrap, o.NoWrap)
				options.StreamLTrim = setInt(options.StreamLTrim, o.StreamLTrim)
				options.StreamRTrim = setInt(options.StreamRTrim, o.StreamRTrim)
				options.MessageLTrim = setInt(options.MessageLTrim, o.MessageLTrim)
				options.MessageRTrim = setInt(options.MessageRTrim, o.MessageRTrim)
				options.MultiLine = setBool(options.MultiLine, o.MultiLine)
				options.DateFormat = setString(options.DateFormat, o.DateFormat)
				options.Green = setString(options.Green, o.Green)
				options.Yellow = setString(options.Yellow, o.Yellow)
				options.Red = setString(options.Red, o.Red)
			}
		}
	}
}

func setString(val string, saved string) string {
	if saved != "" {
		return saved
	}
	return val
}

func setInt(val int, saved int) int {
	if saved > 0 {
		return saved
	}
	return val
}

func setInt64(val int64, saved int64) int64 {
	if saved > 0 {
		return saved
	}
	return val
}

func setBool(val bool, saved bool) bool {
	if saved {
		return true
	}
	return val
}
