package config

import (
	"os"
	"strconv"

	"github.com/mmcquillan/lawsg/util"
)

// EnvVars - Sets the value if the environment variable exists
func EnvVars(options *Options) {

	// filter options
	options.Filter = setEnvString(options.Filter, "LAWSG_FILTER")
	options.Stream = setEnvString(options.Stream, "LAWSG_STREAM")
	options.StartTime = setEnvDate(options.StartTime, "LAWSG_START_TIME")
	options.EndTime = setEnvDate(options.EndTime, "LAWSG_END_TIME")
	options.Number = setEnvInt64(options.Number, "LAWSG_NUMBER")
	options.Tail = setEnvBool(options.Tail, "LAWSG_TAIL")

	// display options
	options.TimeZone = setEnvBool(options.TimeZone, "LAWSG_TIMEZONE")
	options.Spacing = setEnvBool(options.Spacing, "LAWSG_SPACING")
	options.NoGroup = setEnvBool(options.NoGroup, "LAWSG_NO_GROUP")
	options.NoStream = setEnvBool(options.NoStream, "LAWSG_NO_STREAM")
	options.NoTime = setEnvBool(options.NoTime, "LAWSG_NO_TIME")
	options.NoColor = setEnvBool(options.NoColor, "LAWSG_NO_COLOR")
	options.NoWrap = setEnvBool(options.NoWrap, "LAWSG_NO_WRAP")
	options.StreamLTrim = setEnvInt(options.StreamLTrim, "LAWSG_STREAM_LTRIM")
	options.StreamRTrim = setEnvInt(options.StreamRTrim, "LAWSG_STREAM_RTRIM")
	options.MessageLTrim = setEnvInt(options.MessageLTrim, "LAWSG_MESSAGE_LTRIM")
	options.MessageRTrim = setEnvInt(options.MessageRTrim, "LAWSG_MESSAGE_RTRIM")
	options.MultiLine = setEnvBool(options.MultiLine, "LAWSG_MULTI_LINE")
	options.DateFormat = setEnvString(options.DateFormat, "LAWSG_DATE_FORMAT")
	options.Green = setEnvString(options.Green, "LAWSG_GREEN")
	options.Yellow = setEnvString(options.Yellow, "LAWSG_YELLOW")
	options.Red = setEnvString(options.Red, "LAWSG_RED")

	// advanced options
	options.Command = setEnvString(options.Command, "LAWSG_COMMAND")
	options.Group = setEnvString(options.Group, "LAWSG_GROUP")
	options.Chunk = setEnvInt64(options.Chunk, "LAWSG_CHUNK")
	options.Refresh = setEnvInt(options.Refresh, "LAWSG_REFRESH")
	options.Region = setEnvString(options.Region, "LAWSG_REGION")
	options.Env = setEnvString(options.Env, "LAWSG_ENV")
	options.Stats = setEnvBool(options.Stats, "LAWSG_STATS")
	options.Debug = setEnvBool(options.Debug, "LAWSG_DEBUG")
}

func setEnvString(val string, env string) string {
	if os.Getenv(env) != "" {
		return os.Getenv(env)
	}
	return val
}

func setEnvDate(val int64, env string) int64 {
	if os.Getenv(env) != "" {
		return util.ParseDate(os.Getenv(env))
	}
	return val
}

func setEnvInt(val int, env string) int {
	if os.Getenv(env) != "" {
		if v, err := strconv.ParseInt(os.Getenv(env), 10, 32); err == nil {
			return int(v)
		}
	}
	return val
}

func setEnvInt64(val int64, env string) int64 {
	if os.Getenv(env) != "" {
		if v, err := strconv.ParseInt(os.Getenv(env), 10, 64); err == nil {
			return v
		}
	}
	return val
}

func setEnvBool(val bool, env string) bool {
	if os.Getenv(env) != "" {
		if v, err := strconv.ParseBool(os.Getenv(env)); err == nil {
			return v
		}
	}
	return val
}
