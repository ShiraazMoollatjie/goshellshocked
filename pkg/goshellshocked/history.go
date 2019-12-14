package goshellshocked

import (
	"bufio"
	"os"
	"strings"
)

type bashParser struct{}

func (p bashParser) parse(line string) string {
	return line
}

type zshParser struct{}

func (p zshParser) parse(line string) string {
	li := strings.Split(line, ";")
	if len(li) < 2 {
		return ""
	}

	return li[1]
}

type fishParser struct{}

func (f fishParser) parse(line string) string {
	li := strings.Split(line, ": ")
	if len(li) < 2 {
		return ""
	}

	if !strings.HasPrefix(li[0], "- cmd") {
		return ""
	}

	return li[1]
}

func parse(filename, line string) string {
	if strings.Contains(filename, "zsh") {
		return zshParser{}.parse(line)
	} else if strings.Contains(filename, "fish") {
		return fishParser{}.parse(line)
	}

	return bashParser{}.parse(line)
}

// ProcessHistoryFile parses and returns the commands of the provided history file.
func ProcessHistoryFile(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	sc := bufio.NewScanner(f)
	var res []string
	for sc.Scan() {
		w := parse(filename, sc.Text())

		if isExclusion(w) {
			continue
		}

		res = append(res, w)
	}

	return res, nil
}
