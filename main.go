package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var exclusions = flag.String("exclude", "", "A comma separated list of commands. Performs an exact match for each provided word.")

func main() {
	flag.Parse()
	h := os.Getenv("HOME")

	fl, err := ioutil.ReadDir(h)
	if err != nil {
		log.Fatalf("Cannot read from directory %v. Error: %v", h, err)
	}

	wl := map[string]int{}
	for _, f := range fl {
		if strings.Contains(f.Name(), "_history") {
			log.Printf("Found history file: %v", f.Name())
			err := processHistoryFile(filepath.Join(h, f.Name()), wl)
			if err != nil {
				log.Fatalf("Cannot read history file %v. Error: %v", f, err)
			}
		}
	}

	err = buildWorldCloud(wl)
	if err != nil {
		log.Fatalf("Cannot build word cloud. Error: %v", err)
	}
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

func processHistoryFile(filename string, wordList map[string]int) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		w := parse(filename, sc.Text())

		if isExclusion(w) {
			continue
		}

		_, ok := wordList[w]
		if !ok {
			wordList[w] = 1
		} else {
			wordList[w]++
		}
	}

	return nil
}
