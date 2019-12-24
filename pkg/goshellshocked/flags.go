package goshellshocked

import (
	"flag"
	"strings"
)

var exclusions = flag.String("exclude", "", "A comma separated list of commands. Performs an exact match for each provided word.")
var ignore = flag.Int("ignore", 1, "Anything frequency less than the ignored amount should be excluded.")

func isExclusion(command string) bool {
	if command == "" {
		return true
	}

	for _, e := range getExclusions() {
		if command == e {
			return true
		}
	}

	return false
}

func getExclusions() []string {
	return strings.Split(*exclusions, ",")
}
