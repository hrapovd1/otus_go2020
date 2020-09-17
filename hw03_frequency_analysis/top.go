package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"strings"
)

func Top10(inStr string) []string {
	// Прооверка на пустую строку
	if len(inStr) == 0 {
		return []string{}
	}

	// Структура для слова и его счетчика
	type wordRate struct {
		Word  string
		Count int
	}

	words := make([]wordRate, 1)                   // Слайс словесных структур
	var outStr, tmpWords []string                  // Переменные используемые в обработке
	tmpWordsMap := make(map[string]int)            // Словарь слов и их количества
	tmpLines := strings.Split(inStr, string('\n')) // Разбиваю текст на строки

	// Разделяю строки на слова во временный слайс
	for _, sentence := range tmpLines {
		for _, word := range strings.Split(sentence, string(' ')) {
			if len(word) > 0 {
				tmpWords = append(tmpWords, strings.TrimSpace(word))
			}
		}
	}

	// Собираю слова в словарь с одновременным подсчетом одинаковых
	for i := 0; i < len(tmpWords); i++ {
		if len(tmpWords[i]) > 0 {
			tmpWordsMap[tmpWords[i]]++
		}
	}

	// Преобразую словарь слов в слайс структуры слово-счетчик для сортировки
	for word, count := range tmpWordsMap {
		words = append(words, wordRate{Word: word, Count: count})
	}
	// Сортировка слайса структур по счетчику слов от большего к меньшему
	sort.Slice(words, func(i, j int) bool {
		return words[i].Count > words[j].Count
	})

	// Беру первые 10 слов из сортированного по счетчику слайса
	// и возврощаю полученный результат
	for i := 0; i < 10; i++ {
		outStr = append(outStr, words[i].Word)
	}
	return outStr
}
