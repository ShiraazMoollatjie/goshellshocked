package goshellshocked

import (
	"reflect"
	"testing"
)

func setIgnoreForTesting(t *testing.T, minOccurence int) func() {
	return func() {
		old := ignore
		*ignore = minOccurence
		defer func() {
			*ignore = *old
		}()
	}
}

func TestToCommands(t *testing.T) {
	testCases := []struct {
		desc             string
		shellHistory     []string
		minOccurrences   int
		expectedCommands []string
	}{
		{
			desc: "basic test",
			shellHistory: []string{
				"ls",
				"ls",
			},
			minOccurrences: 1,
			expectedCommands: []string{
				"ls",
			},
		},
		{
			desc: "return no commands",
			shellHistory: []string{
				"ls",
				"ls",
			},
			minOccurrences:   3,
			expectedCommands: nil,
		},
		{
			desc: "return a mix of commands in descending order",
			shellHistory: []string{
				"git branch -v",
				"ls",
				"ls",
			},
			minOccurrences: 1,
			expectedCommands: []string{
				"ls",
				"git branch -v",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			setIgnoreForTesting(t, tC.minOccurrences)()
			c := ToCommands(tC.shellHistory)
			if !reflect.DeepEqual(c.data, tC.expectedCommands) {
				t.Errorf("incorrect commands generated for shellHistory %v. expected %v, generated %v", tC.shellHistory, tC.expectedCommands, c.data)
			}
		})
	}
}
