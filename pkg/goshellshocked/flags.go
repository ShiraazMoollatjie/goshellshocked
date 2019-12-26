package goshellshocked

import (
	"flag"
	"os"
	"strings"
)

var exclusions = flag.String("exclude", "", "A comma separated list of commands. Performs an exact match for each provided word.")
var ignore = flag.Int("ignore", 3, "Anything frequency less than the ignored amount (exclusive) should be excluded.")
var output = flag.String("output", "console", "The output mode for the commands. Defaults to stdout.")
var outputDir = flag.String("outputDir", "", "The output directory to use.")

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

func getOutput() string {
	return *output
}

func getOutputDir() (string, error) {
	if *outputDir == "" {
		return os.Getwd()
	}

	return *outputDir, nil
}
