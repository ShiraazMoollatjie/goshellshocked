package goshellshocked

func ToFrequencyMap(wl []string) map[string]int {
	res := map[string]int{}

	for _, w := range wl {
		_, ok := res[w]
		if !ok {
			res[w] = 1
		} else {
			res[w]++
		}
	}

	freq := map[string]int{}
	for k, v := range res {
		if v > *minOccurrences {
			freq[k] = v
		}
	}

	return freq
}
