package goshellshocked

import "sort"

// Commands is a container for a list of commands along with their frequencies..
type Commands struct {
	data        []string
	frequencies map[string]int
}

func (s Commands) Len() int { return len(s.data) }

func (s Commands) Less(i, j int) bool {
	return s.frequencies[s.data[i]] > s.frequencies[s.data[j]]
}

func (s Commands) Swap(i, j int) {
	s.data[j], s.data[i] = s.data[i], s.data[j]
}

// GetFrequency returns the frequency of the provided command.
func (s Commands) GetFrequency(command string) int {
	return s.frequencies[command]
}

// GetData returns the underlying slice of commands.
func (s Commands) GetData() []string {
	return s.data
}

// FilterFrequencies returns the Commands that have a higher or equal frequency count than the one provided.
func (s Commands) FilterFrequencies(frequency int) Commands {
	var resData []string
	resFreq := map[string]int{}
	for k, v := range s.frequencies {
		if v > frequency {
			resData = append(resData, k)
			resFreq[k] = v
		}
	}

	res := Commands{
		data:        resData,
		frequencies: resFreq,
	}
	sort.Sort(res)

	return res
}

// ToCommands converts a slice of terminal commands into a Commands type.
func ToCommands(wl []string) Commands {
	freq := map[string]int{}

	for _, w := range wl {
		_, ok := freq[w]
		if !ok {
			freq[w] = 1
		} else {
			freq[w]++
		}
	}

	commands := Commands{
		data:        wl,
		frequencies: freq,
	}

	return commands.FilterFrequencies(*ignore)
}
