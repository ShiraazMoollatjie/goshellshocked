package main

import (
	"bufio"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/psykhi/wordclouds"
)

func main() {
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

func buildWorldCloud(wordList map[string]int) error {

	colors := make([]color.Color, 0)
	for _, c := range []color.RGBA{
		{0x1b, 0x1b, 0x1b, 0xff},
		{0x48, 0x48, 0x4B, 0xff},
		{0x59, 0x3a, 0xee, 0xff},
		{0x65, 0xCD, 0xFA, 0xff},
		{0x70, 0xD6, 0xBF, 0xff},
	} {
		colors = append(colors, c)
	}

	w := wordclouds.NewWordcloud(wordList,
		wordclouds.Height(2048),
		wordclouds.Width(2048),
		wordclouds.FontFile("Roboto-Regular.ttf"),
		wordclouds.FontMaxSize(700),
		wordclouds.FontMinSize(10),
		wordclouds.Colors(colors),
	)

	img := w.Draw()

	outputFile, err := os.Create("wc.png")
	if err != nil {
		// Handle error
		return err
	}
	defer outputFile.Close()

	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(outputFile, img)

	return nil
}

func processHistoryFile(filename string, wordList map[string]int) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		w := parseHistoryLine(filename, sc.Text())

		_, ok := wordList[w]
		if !ok {
			wordList[w] = 1
		} else {
			wordList[w]++
		}
	}

	return nil
}

func parseHistoryLine(filename, line string) string {
	if !strings.Contains(filename, "zsh") {
		return line
	}

	li := strings.Split(line, ";")
	if len(li) < 2 {
		return ""
	}

	return li[1]
}
