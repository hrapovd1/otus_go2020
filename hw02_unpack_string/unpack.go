package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
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

	strRune := []rune(str)

	// Если первая цифра, то выйти с ошибкой
	if unicode.IsDigit(strRune[0]) {
		return "", ErrInvalidString
	}

	var strOut strings.Builder // Объект результирующей строки
	isDouble := false          // Флаг для обработки эскейп последовательностей вида: '\n'
	const backSlash = '\\'     // Оптимизаци линтера goconst

	// Основной цикл перебора символов
	for pos := 0; pos < len(strRune); pos++ {
		// Для удобства понимания и работы с кодом определяю переменные предыдущего и текущего символов
		var prevChar, currChar rune
		if pos > 0 {
			prevChar, currChar = strRune[pos-1], strRune[pos]
		} else {
			prevChar, currChar = 0, strRune[pos]
		}

		// Проверка на последний символ
		notLastChar := true // Установка флага, что это не последний символ
		if pos == len(strRune)-1 {
			notLastChar = false
		}

		switch {
		// Обрабатываю первый символ
		case pos == 0:
			switch {
			case notLastChar && unicode.IsDigit(strRune[pos+1]): // Если следующий символ цифра, то обработать на следующем цикле
				break
			case currChar == backSlash: // Если эскейп экранирование, то обработать на следующем цикле
				if notLastChar {
					break
				} else {
					strOut.WriteRune(currChar)
				}
			default: // Если не цифра, эскейп экранирование и следующий символ не цифра, то записываем в выходную строку
				strOut.WriteRune(currChar)
			}

		// Обработка символов кроме первого
		case pos != 0:
			switch {
			case unicode.IsDigit(currChar):
				// Проверка что это не число
				if notLastChar && unicode.IsDigit(strRune[pos+1]) || unicode.IsDigit(strRune[pos-1]) {
					//	return "", errors.New("wrong string, it can be only digit, not numeric")
					return "", ErrInvalidString
				}

				// Преобразование цифры из строки в целое
				currDigit, err := strconv.Atoi(string(currChar))
				if err != nil {
					return "", err
				}

				writedChar := string(prevChar)
				// Если предыдущий символ из эскейп последовательности
				if isDouble {
					writedChar = string(strRune[pos-2 : pos])
				}
				for count := currDigit; count > 0; count-- {
					strOut.WriteString(writedChar)
				}
			case currChar == backSlash:
				if prevChar == backSlash {
					strOut.WriteString(string(prevChar))
				}
			default:
				if prevChar == backSlash {
					// Проверка на эскейп последовательность
					if notLastChar && unicode.IsDigit(strRune[pos+1]) {
						isDouble = true
						break
					}
					strOut.WriteString(string(prevChar))
				}
				if notLastChar && unicode.IsDigit(strRune[pos+1]) {
					break
				}
				strOut.WriteString(string(currChar))
			}
		}
	}
	return strOut.String(), nil
}
