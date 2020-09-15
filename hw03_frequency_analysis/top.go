package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"strings"
)

func Top10(inStr string) []string {
	if len(inStr) == 0 {
		return []string{}
	}
	type wordRate struct {
		Word  string
		Count int
	}
	words := make([]wordRate, 1)
	var outStr, tmpWords []string
	tmpWordsMap := make(map[string]int)
	tmpLines := strings.Split(strings.ToLower(inStr), "\n")

	for _, sentence := range tmpLines {
		for _, word := range strings.Split(sentence, " ") {
			if len(word) > 0 {
				tmpWords = append(tmpWords, strings.TrimSpace(word))
			}
		}
	}

	for i := 0; i < len(tmpWords); i++ {
		/*
			for _, word := range strings.Split(tmpWords[i], ".") {
				if len(word) > 0 {
					tmpWordsMap[word]++
				}
			}
		*/
		if len(tmpWords[i]) > 0 {
			tmpWordsMap[tmpWords[i]]++
		}
	}

	//fmt.Println("{------- start -----------}")
	for word, count := range tmpWordsMap {
		words = append(words, wordRate{Word: word, Count: count})
	}
	sort.Slice(words, func(i, j int) bool {
		return words[i].Count > words[j].Count
	})
	for i := 0; i < 10; i++ {
		//	fmt.Printf("Word: %s  ; Count: %d\n", words[i].Word, words[i].Count)
		outStr = append(outStr, words[i].Word)
	}
	return outStr
}
