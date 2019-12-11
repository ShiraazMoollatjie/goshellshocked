package goshellshocked

import (
	"flag"
	"strings"
)

var exclusions = flag.String("exclude", "", "A comma separated list of commands. Performs an exact match for each provided word.")
var minOccurrences = flag.Int("minCount", 1, "The minimum frequency count for the command to be included.")

func isExclusion(term string) bool {
	for _, e := range getExclusions() {
		if term == e {
			return true
		}
	}

	return false
}

func getExclusions() []string {
	return strings.Split(*exclusions, ",")
}
