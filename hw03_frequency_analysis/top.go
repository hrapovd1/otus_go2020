package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"strings"
	"unicode"
)

func Top10(inStr string) []string {
	// Прооверка на пустую строку
	if len(inStr) == 0 {
		return nil
	}

	// Структура для слова и его счетчика
	type wordRate struct {
		Word  string
		Count int
	}

	words := []wordRate{}                                           // Слайс словесных структур
	outStr, tmpWords := []string{}, []string{}                      // Переменные используемые в обработке
	var rateCount int                                               // Количество возвращаемых первых слов
	tmpWordsMap := make(map[string]int)                             // Словарь слов и их количества
	tmpLines := strings.Split(strings.ToLower(inStr), string('\n')) // Разбиваю текст на строки и делаю все буквы строчными

	// Разделяю строки на слова во временный слайс
	for _, sentence := range tmpLines {
		tmpWords = append(tmpWords, strings.FieldsFunc(sentence, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsPunct(r)
		})...)
	}

	// Собираю слова в словарь с одновременным подсчетом одинаковых
	for i := 0; i < len(tmpWords); i++ {
		cleanedWord := cleanPunct(tmpWords[i]) // Чищу слово от пунктуации
		if len(cleanedWord) > 0 {
			tmpWordsMap[cleanedWord]++
		}
	}

	// Преобразую словарь слов в слайс структуры слово-счетчик для сортировки
	words = make([]wordRate, 0, len(tmpWordsMap))
	for word, count := range tmpWordsMap {
		words = append(words, wordRate{Word: word, Count: count})
	}
	// Сортировка слайса структур по счетчику слов от большего к меньшему
	sort.Slice(words, func(i, j int) bool {
		return words[i].Count > words[j].Count
	})

	// Беру первые 10 слов из сортированного по счетчику слайса
	// и возврощаю полученный результат
	if len(words) < 11 {
		rateCount = len(words)
	} else {
		rateCount = 10
	}
	for i := 0; i < rateCount; i++ {
		outStr = append(outStr, words[i].Word)
	}
	return outStr
}

// Функция очистки слов от пунктуации.
func cleanPunct(s string) string {
	str := s
	outStr := ""
	for _, symbol := range str {
		if !unicode.IsPunct(symbol) {
			outStr += string(symbol)
		}
	}
	return outStr
}
