package util

import (
	"strconv"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

//ParseDate - Parse date time from a variety of formats
func ParseDate(input string) int64 {
	input = strings.TrimSpace(strings.ToLower(input))
	if input == "n" || input == "now" {
		return time.Now().Unix() * 1000
	}
	if strings.Contains(input, "ago") {
		input = strings.TrimSpace(strings.Replace(input, "ago", "", 1))
		d, err := time.ParseDuration(input)
		if err == nil {
			return time.Now().Add(-d).Unix() * 1000
		}
		if strings.Contains(input, "d") {
			input = strings.TrimSpace(strings.Replace(input, "d", "", 1))
			i, err := strconv.Atoi(input)
			if err != nil {
				return 0
			}
			return time.Now().AddDate(0, 0, -i).Unix() * 1000
		}
		if strings.Contains(input, "w") {
			input = strings.TrimSpace(strings.Replace(input, "w", "", 1))
			i, err := strconv.Atoi(input)
			if err != nil {
				return 0
			}
			return time.Now().AddDate(0, 0, -(i*7)).Unix() * 1000
		}
		if strings.Contains(input, "y") {
			input = strings.TrimSpace(strings.Replace(input, "y", "", 1))
			i, err := strconv.Atoi(input)
			if err != nil {
				return 0
			}
			return time.Now().AddDate(-i, 0, 0).Unix() * 1000
		}
	}
	t, err := dateparse.ParseIn(input, time.UTC)
	if err != nil {
		return 0
	}
	return t.Unix() * 1000
}
