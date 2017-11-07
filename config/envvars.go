package config

import (
	"os"
	"strconv"

	"github.com/mmcquillan/lawsg/util"
)

// EnvVars - Sets the value if the environment variable exists
func EnvVars(options *Options) {
	options.Command = setString(options.Command, "LAWSG_COMMAND")
	options.Group = setString(options.Group, "LAWSG_GROUP")
	options.Filter = setString(options.Filter, "LAWSG_FILTER")
	options.Stream = setString(options.Stream, "LAWSG_STREAM")
	options.StartTime = setDate(options.StartTime, "LAWSG_START_TIME")
	options.EndTime = setDate(options.EndTime, "LAWSG_END_TIME")
	options.Number = setInt64(options.Number, "LAWSG_NUMBER")
	options.Chunk = setInt64(options.Chunk, "LAWSG_CHUNK")
	options.Tail = setBool(options.Tail, "LAWSG_TAIL")
	options.TimeZone = setBool(options.TimeZone, "LAWSG_TIMEZONE")
	options.Spacing = setBool(options.Spacing, "LAWSG_SPACING")
	options.NoGroup = setBool(options.NoGroup, "LAWSG_NO_GROUP")
	options.NoStream = setBool(options.NoStream, "LAWSG_NO_STREAM")
	options.NoTime = setBool(options.NoTime, "LAWSG_NO_TIME")
	options.NoColor = setBool(options.NoColor, "LAWSG_NO_COLOR")
	options.NoWrap = setBool(options.NoWrap, "LAWSG_NO_WRAP")
	options.TrimLeft = setInt(options.TrimLeft, "LAWSG_TRIM_LEFT")
	options.DateFormat = setString(options.DateFormat, "LAWSG_DATE_FORMAT")
	options.Green = setString(options.Green, "LAWSG_GREEN")
	options.Yellow = setString(options.Yellow, "LAWSG_YELLOW")
	options.Red = setString(options.Red, "LAWSG_RED")
	options.Refresh = setInt(options.Refresh, "LAWSG_REFRESH")
	options.Cache = setBool(options.Cache, "LAWSG_CACHE")
	options.CacheDir = setString(options.CacheDir, "LAWSG_CACHE_DIR")
	options.Debug = setBool(options.Debug, "LAWSG_DEBUG")
}

func setString(val string, env string) string {
	if os.Getenv(env) != "" {
		return os.Getenv(env)
	}
	return val
}

func setDate(val int64, env string) int64 {
	if os.Getenv(env) != "" {
		return util.ParseDate(os.Getenv(env))
	}
	return val
}

func setInt(val int, env string) int {
	if os.Getenv(env) != "" {
		if v, err := strconv.ParseInt(os.Getenv(env), 10, 32); err == nil {
			return int(v)
		}
	}
	return val
}

func setInt64(val int64, env string) int64 {
	if os.Getenv(env) != "" {
		if v, err := strconv.ParseInt(os.Getenv(env), 10, 64); err == nil {
			return v
		}
	}
	return val
}

func setBool(val bool, env string) bool {
	if os.Getenv(env) != "" {
		if v, err := strconv.ParseBool(os.Getenv(env)); err == nil {
			return v
		}
	}
	return val
}
