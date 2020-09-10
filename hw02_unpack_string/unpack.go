package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	// Проверка на пустую строку
	if len(str) == 0 {
		return "", nil
	}

	strRune := []rune(str) // Разделяю строку на руны

	// Проверка на первую цифру
	if unicode.IsDigit(strRune[0]) {
		return "", ErrInvalidString
	}

	var strOut strings.Builder // Объект результирующей строки
	isPrevDigit := false          // Флаг нахождения цифры в предыдущем символе
	isDoubleDigit := false        // Флаг нахождения последовательности 2х цифр
	isPrevSlash := false // Флаг обработки двойных слешей
	isLastSymbol := false         // Флаг последнего символа
	const backSlash = '\\'     // Оптимизаци линтера goconst

	// Основной цикл перебора символов
	for pos := len(strRune)-1; pos >= 0 ; pos-- {
		// Для удобства понимания и работы с кодом определяю переменные следующего и текущего символов
		var nextChar, currChar rune
		if pos != 0 {
			nextChar, currChar = strRune[pos-1], strRune[pos]
		} else {
			nextChar, currChar = 0, strRune[pos]
		}
		if pos == len(strRune)-1{
			isLastSymbol = true
		}

		switch {
		case unicode.IsDigit(currChar):
			if isPrevDigit {
				isPrevDigit = false
				isDoubleDigit = true
			} else {
				isPrevDigit = true
			}
		case currChar == backSlash:
			switch {
			case isLastSymbol:
				strOut.WriteRune(currChar)
			case isPrevDigit:
				strOut.WriteRune(strRune[pos + 1])
				isPrevDigit = false
			case isDoubleDigit:
				writeReapetedSymbol(string(strRune[pos + 1]), strRune[pos + 2], &strOut)
				isDoubleDigit = false
			case isPrevSlash:
				if pos+2
			}
		case !unicode.IsDigit(currChar):
		}
	}
	return strOut.String(), nil
}

// Функция распаковки символа в выходную строку, принимает
// символ или комбинацию, например \p, которые записываем = r
// сколько раз повторить = r2
// объект выходной строки s
func writeReapetedSymbol(r string, r2 rune, s *strings.Builder) {
	i, err := strconv.Atoi(string(r2))
	if err != nil{
		log.Fatalf("Atoi error %v", err)
	}
	for ;i > 0; i--{
		s.WriteString(r)
	}
}
