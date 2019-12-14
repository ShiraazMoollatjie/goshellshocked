package goshellshocked

import "testing"

func TestFishParsing(t *testing.T) {
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

		})
	}
}
