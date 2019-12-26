package goshellshocked

import (
	"flag"
	"strings"
)

var exclusions = flag.String("exclude", "", "A comma separated list of commands. Performs an exact match for each provided word.")
var ignore = flag.Int("ignore", 3, "Anything frequency less than the ignored amount (exclusive) should be excluded.")
var output = flag.String("output", "console", "The output mode for the commands. Defaults to stdout.")

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

// GetPrintOption returns the PrintOtion based on the option flag
func GetPrintOption() PrintOption {
	switch *output {
	case "json":
		return PrintOptionJSON
	case "yaml":
		return PrintOptionYAML
	default:
		return PrintOptionStdout
	}
}
