package goshellshocked

import "strings"

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

func Parse(filename, line string) string {
	if strings.Contains(filename, "zsh") {
		return newZshParser().parse(line)
	}

	return newBashParser().parse(line)
}
