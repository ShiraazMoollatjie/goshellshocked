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

func newBashParser() *bashParser {
	return &bashParser{}
}

type zshParser struct{}

func newZshParser() *zshParser {
	return &zshParser{}
}

func (p zshParser) parse(line string) string {
	li := strings.Split(line, ";")
	if len(li) < 2 {
		return ""
	}

	return li[1]
}

func parse(filename, line string) string {
	if strings.Contains(filename, "zsh") {
		return newZshParser().parse(line)
	}

	return newBashParser().parse(line)
}

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
