package util

import (
	"regexp"
)

// CountFormat - counts how many formatted characters in the string
func CountFormat(format string) int {
	reg := regexp.MustCompile("(\\033|\\027)\\[[0-9]*m")
	res := reg.FindAllStringIndex(format, -1)
	return len(res)
}
