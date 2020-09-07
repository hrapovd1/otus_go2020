package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

/*
Не знаю как использовать
var ErrInvalidString = errors.New("invalid string")
*/

func Unpack(str string) (string, error) {
	// Проверка на пустую строку
	if len(str) == 0 {
		return "", nil
	}

	var strOut strings.Builder // Объект результирующей строки
	isDouble := false          // Флаг для обработки эскейп последовательностей вида: '\n'

	// Основной цикл перебора символов
	for pos := 0; pos < len(str); pos++ {

		switch {
		// Проверка первого символа на цифру и эскейп экран
		case pos == 0:
			switch {
			case unicode.IsDigit(rune(str[pos])): // Если первая цифра, то выйти с ошибкой
				return "", errors.New("wrong string, it can't start from digit")
			case unicode.IsDigit(rune(str[pos+1])): // Если следующий символ цифра, то обработать на следующем цикле
				break
			case string(str[pos]) == `\`: // Если эскейп экран, то обработать на следующем цикле
				break
			default: // Если не цифра, эскейп экран и следующий символ не цифра, то записываем в выходную строку
				strOut.WriteString(string(str[pos]))
			}

		// Обработка символов кроме первого
		case pos != 0:
			notLastChar := true    // Установка флага, что это не последний символ
			if pos == len(str)-1 { // Проверка на последний символ
				notLastChar = false
			}
			// Для удобства понимания и работы с кодом определяю переменные предыдущего и текущего символов
			prevChar, currChar := str[pos-1], str[pos]

			switch {
			case unicode.IsDigit(rune(currChar)):
				// Проверка что это не число
				if notLastChar && unicode.IsDigit(rune(str[pos+1])) || unicode.IsDigit(rune(str[pos-1])) {
					return "", errors.New("wrong string, it can be only digit, not numeric")
				}
				// Преобразование цифры из строки в целое
				currDigit, err := strconv.Atoi(string(currChar))
				if err != nil {
					return "", err
				}

				switch {
				case currDigit == 0:
					break
				case currDigit > 0:
					writedChar := string(prevChar)
					// Если предыдущий символ из эскейп последовательности
					if isDouble {
						writedChar = str[pos-2 : pos]
					}
					for count := currDigit; count > 0; count-- {
						strOut.WriteString(writedChar)
					}
				}
			case string(currChar) == `\`:
				if string(prevChar) == `\` {
					strOut.WriteString(string(prevChar))
				}
				break
			default:
				if string(prevChar) == `\` {
					// Проверка на эскейп последовательность
					if notLastChar && unicode.IsDigit(rune(str[pos+1])) {
						isDouble = true
						break
					}
					strOut.WriteString(string(prevChar))
				}
				if notLastChar && unicode.IsDigit(rune(str[pos+1])) {
					break
				}
				strOut.WriteString(string(currChar))
			}
		}
	}
	return strOut.String(), nil
}
