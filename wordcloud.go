package main

import (
	"image/color"
	"image/png"
	"os"

	"github.com/psykhi/wordclouds"
)

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
		wordclouds.FontMaxSize(100),
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
