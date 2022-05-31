package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

type words []wordCount

func (s words) Len() int {
	return len(s)
}

func (s words) Less(i, j int) bool {
	if s[i].count == s[j].count {
		return strings.Compare(s[i].word, s[j].word) < 0
	}
	return s[i].count > s[j].count
}

func (s words) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func Top10(text string) []string {
	if text == "" {
		return nil
	}
	wordsInText := strings.Fields(text)
	wordC := make(map[string]int)
	for _, word := range wordsInText {
		wordC[word]++
	}
	w := make(words, 0)
	for k, v := range wordC {
		w = append(w, wordCount{word: k, count: v})
	}
	sort.Sort(w)
	topWordsSize := func() int {
		if len(w) > 10 {
			return 10
		}
		return len(w)
	}()
	topWords := make([]string, 0)
	for i := 0; i < topWordsSize; i++ {
		topWords = append(topWords, w[i].word)
	}
	return topWords
}
