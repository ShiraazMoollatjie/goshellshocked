package goshellshocked

import (
	"testing"
)

func TestFishParser(t *testing.T) {
	testCases := []struct {
		desc    string
		command string
		result  string
	}{
		{"Basic command", "- cmd: git checkout master", "git checkout master"},
		{"when line, should return a blank string", "  when: 1575976962", ""},
		{"junk test, should return a blank string", "junk", ""},
		{"empty string, should return a blank string", "", ""},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := parseFish(tC.command)
			if res != tC.result {
				t.Fatalf("expected %v, got %v", tC.result, res)
			}
		})
	}
}

func TestZSHParser(t *testing.T) {
	testCases := []struct {
		desc    string
		command string
		result  string
	}{
		{"Basic command", "576359202:0;cat ~/.zsh_history", "cat ~/.zsh_history"},
		{"when line, should return a blank string", "576359202:0", ""},
		{"junk test, should return a blank string", "junk", ""},
		{"empty string, should return a blank string", "", ""},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := parseZsh(tC.command)
			if res != tC.result {
				t.Fatalf("expected %v, got %v", tC.result, res)
			}
		})
	}
}

func TestBashParser(t *testing.T) {
	testCases := []struct {
		desc    string
		command string
		result  string
	}{
		{"Basic command", "history | grep pacman", "history | grep pacman"},
		{"junk test, should return the same string", "junk", "junk"},
		{"empty string, should return a blank string", "", ""},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := parseBash(tC.command)
			if res != tC.result {
				t.Fatalf("expected %v, got %v", tC.result, res)
			}
		})
	}
}
