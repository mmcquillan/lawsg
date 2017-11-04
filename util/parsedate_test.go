package util

import (
	"testing"
)

func TestTime(t *testing.T) {
	checkTime("n", t)
	checkTime("now", t)
	checkTime("1s ago", t)
	checkTime("1m ago", t)
	checkTime("1h ago", t)
	checkTime("1d ago", t)
	checkTime("1w ago", t)
	checkTime("1y ago", t)
	checkTime("2017-11-01 12:30:45", t)
}

func checkTime(chk string, t *testing.T) {
	if ParseDate(chk) == 0 {
		t.Fatalf("Failed DateTime Parsing for: " + chk)
	}
}
