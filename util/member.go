package util

import (
	"strings"
)

// Member - Determine if a needle is in a comma delimited haystack
func Member(needle string, haystack string) bool {
	for _, h := range strings.Split(haystack, ",") {
		if strings.ToUpper(needle) == strings.ToUpper(h) {
			return true
		}
	}
	return false
}
