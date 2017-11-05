package joda

import (
	"strings"
)

func Format(joda string) (goda string) {

	// mapping of values
	damap := map[string]string{
		"y":    "06",
		"Y":    "2006",
		"dd":   "02",
		"d":    "2",
		"hh":   "03",
		"h":    "3",
		"H":    "15",
		"mm":   "04",
		"m":    "4",
		"ss":   "05",
		"s":    "5",
		"S":    ".0",
		"Z":    "-0700",
		"a":    "pm",
		"z":    "MST",
		"MMMM": "January",
		"MMM":  "Jan",
		"MM":   "01",
		"M":    "1",
		"EE":   "Monday",
		"E":    "Mon",
	}

	// order of mapping
	order := []string{
		"y",
		"Y",
		"dd",
		"d",
		"hh",
		"h",
		"H",
		"mm",
		"m",
		"ss",
		"s",
		"S",
		"Z",
		"a",
		"z",
		"MMMM",
		"MMM",
		"MM",
		"M",
		"EE",
		"E",
	}

	// run through and replace
	goda = joda
	for i, _ := range order {
		goda = strings.Replace(goda, order[i], damap[order[i]], -1)
	}
	return goda

}
