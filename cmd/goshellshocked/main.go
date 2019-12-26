package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ShiraazMoollatjie/goshellshocked/pkg/goshellshocked"
)

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
			w, err := goshellshocked.ProcessHistoryFile(filepath.Join(h, f.Name()))
			if err != nil {
				log.Fatalf("Cannot read history file %v. Error: %v", f, err)
			}
			wl = append(wl, w...)
		}
	}

	cl := goshellshocked.ToCommands(wl)
	goshellshocked.Write(cl)
}
