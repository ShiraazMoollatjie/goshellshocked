package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ShiraazMoollatjie/goshellshocked"
)

var exclusions = flag.String("exclude", "", "A comma separated list of commands. Performs an exact match for each provided word.")
var minOccurrences = flag.Int("minCount", 1, "The minimum frequency count for the command to be included.")

func main() {
	flag.Parse()
	h := os.Getenv("HOME")

	fl, err := ioutil.ReadDir(h)
	if err != nil {
		log.Fatalf("Cannot read from directory %v. Error: %v", h, err)
	}

	var wl []string
	for _, f := range fl {
		if strings.Contains(f.Name(), "_history") {
			log.Printf("Found history file: %v", f.Name())
			w, err := processHistoryFile(filepath.Join(h, f.Name()))
			if err != nil {
				log.Fatalf("Cannot read history file %v. Error: %v", f, err)
			}
			wl = append(wl, w...)
		}
	}

	err = goshellshocked.BuildWorldCloud(toFrequencyMap(wl))
	if err != nil {
		log.Fatalf("Cannot build word cloud. Error: %v", err)
	}
}

func toFrequencyMap(wl []string) map[string]int {
	res := map[string]int{}

	for _, w := range wl {
		_, ok := res[w]
		if !ok {
			res[w] = 1
		} else {
			res[w]++
		}
	}

	freq := map[string]int{}
	for k, v := range res {
		if v > *minOccurrences {
			freq[k] = v
		}
	}

	return freq
}

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

func processHistoryFile(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	sc := bufio.NewScanner(f)
	var res []string
	for sc.Scan() {
		w := goshellshocked.Parse(filename, sc.Text())

		if isExclusion(w) {
			continue
		}

		res = append(res, w)
	}

	return res, nil
}
