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
	isPrevDigit := false       // Флаг нахождения цифры в предыдущем символе
	waitSlash := false         // Флаг для комбинации вида \n
	//	isDoubleDigit := false        // Флаг нахождения последовательности 2х цифр
	//	isPrevSlash := false // Флаг обработки двойных слешей
	const backSlash = '\\' // Оптимизаци линтера goconst

	// Основной цикл перебора символов
	for pos := len(strRune) - 1; pos >= 0; pos-- {
		// Для удобства понимания и работы с кодом определяю переменную текущего символа
		currChar := strRune[pos]
		isLastSymbol := false // Флаг последнего символа
		if pos == len(strRune)-1 {
			isLastSymbol = true
		}

		switch {
		case unicode.IsDigit(currChar):
			switch {
			case isPrevDigit:
				return "", ErrInvalidString
			default:
				isPrevDigit = true
			}
		case currChar == backSlash:
			//TODO: додумать разные комбинации
			switch {
			case isLastSymbol:
				strOut.WriteRune(currChar)
			case isPrevDigit:
				writeRepeatedSymbol(string(currChar), strRune[pos+1], &strOut)
				isPrevDigit = false
			case waitSlash:
				writeRepeatedSymbol(string(strRune[pos:pos+2]), strRune[pos+1], &strOut)
				waitSlash = false
				isPrevDigit = false
				/*			case isDoubleDigit:
								writeRepeatedSymbol(string(strRune[pos + 1]), strRune[pos + 2], &strOut)
								isDoubleDigit = false
							case isPrevSlash:
								if pos+2 <= len(strRune)-1 && unicode.IsDigit(strRune[pos+2]){
									writeRepeatedSymbol(string(currChar),strRune[pos+2], &strOut)
									isPrevSlash = false
								} */
			}
		case !unicode.IsDigit(currChar):
			switch {
			case isLastSymbol:
				strOut.WriteRune(currChar)
			case isPrevDigit:
				waitSlash = true
			case waitSlash:
				writeRepeatedSymbol(string(strRune[pos+1]), strRune[pos+2], &strOut)
				waitSlash = false
				isPrevDigit = false
			default:
				strOut.WriteRune(currChar)
			}
		}
	}
	return strOut.String(), nil
}

// Функция распаковки символа в выходную строку, принимает
// символ или комбинацию, например \p, которые записываем = r
// сколько раз повторить = r2
// объект выходной строки s.
func writeRepeatedSymbol(r string, r2 rune, s *strings.Builder) {
	i, err := strconv.Atoi(string(r2))
	if err != nil {
		log.Fatalf("Atoi error %v", err)
	}
	for ; i > 0; i-- {
		s.WriteString(r)
	}
}
