package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var r = regexp.MustCompile(`\w`)

type topWord struct {
	w string
	c int
}

func Top10(input string) []string {
	if len(input) == 0 {
		return make([]string, 0)
	}

	input = r.ReplaceAllString(input, " ")
	words := strings.Fields(input)
	m := make(map[string]int, len(words))

	for _, word := range words {
		m[word]++
	}

	topWords := make([]topWord, len(words))
	mIndex := 0

	for word, count := range m {
		topWords[mIndex] = topWord{word, count}
		mIndex++
	}

	sort.Slice(topWords, func(i, j int) bool {
		if topWords[i].c == topWords[j].c {
			return topWords[i].w < topWords[j].w
		}

		return topWords[i].c > topWords[j].c
	})

	top10 := make([]string, 10)

	for i, tWord := range topWords {
		if i == 10 {
			break
		}

		top10[i] = tWord.w
	}

	return top10
}
