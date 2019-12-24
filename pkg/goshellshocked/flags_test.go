package goshellshocked

import (
	"fmt"
	"testing"
)

func setExclusionForTesting(t *testing.T, exclusion string) func() {
	return func() {
		old := exclusions
		*exclusions = exclusion
		defer func() {
			*exclusions = *old
		}()
	}
}
func TestExclusions(t *testing.T) {
	testCases := []struct {
		desc      string
		exclusion string
		command   string
		result    bool
	}{
		{"basic exclusion", "ls", "ls", true},
		{"no exclusion", "", "ls", false},
		{"case matters", "git", "GIT", false},
		{"a subset of a command is not an exclusion", "git commit", "git commit -am  ", false},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			setExclusionForTesting(t, tC.exclusion)()
			res := isExclusion(tC.command)
			if res != tC.result {
				t.Errorf(fmt.Sprintf("expected %v, got %v for command %v", tC.result, res, tC.exclusion))
			}
		})
	}
}
