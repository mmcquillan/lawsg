package util

import (
	"regexp"
)

func Unformat(format string) (unformat string) {
	reg := regexp.MustCompile("(\\033|\\027)\\[[0-9]*m")
	unformat = reg.ReplaceAllString(format, "")
	return unformat
}
