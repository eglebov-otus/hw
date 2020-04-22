//nolint:golint,stylecheck
package hw03_frequency_analysis

import (
	"regexp"
	"sort"
	"strings"
)

type Word struct {
	Word string
	Freq int
}

const WordsLimit = 10

func Top10(s string) []string {
	// build freq map
	freqMap := map[string]int{}

	r := regexp.MustCompile(`'?\pL[\pL']*(?:-\pL+)*'?`)

	for _, m := range r.FindAllString(s, -1) {
		freqMap[strings.ToLower(m)]++
	}

	words := make([]Word, 0, len(freqMap))

	for word, freq := range freqMap {
		words = append(words, Word{word, freq})
	}

	// sort words by freq
	sort.Slice(words, func(i, j int) bool {
		return words[i].Freq > words[j].Freq
	})

	res := make([]string, 0, len(words))

	for _, word := range words {
		res = append(res, word.Word)
	}

	if len(res) > WordsLimit {
		return res[:WordsLimit]
	}

	return res
}
