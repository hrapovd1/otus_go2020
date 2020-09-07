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

	var strOut strings.Builder // Объект результирующей строки
	isDouble := false          // Флаг для обработки эскейп последовательностей вида: '\n'
	const backSlash = `\`      // Оптимизаци линтера goconst

	// Основной цикл перебора символов
	for pos := 0; pos < len(str); pos++ {
		// Для удобства понимания и работы с кодом определяю переменные предыдущего и текущего символов
		var prevChar, currChar rune
		if pos > 0 {
			prevChar, currChar = rune(str[pos-1]), rune(str[pos])
		} else {
			prevChar, currChar = 0, rune(str[pos])
		}
		notLastChar := true // Установка флага, что это не последний символ
		// Проверка на последний символ
		if pos == len(str)-1 {
			notLastChar = false
		}

		switch {
		// Проверка первого символа на цифру и эскейп экран
		case pos == 0:
			switch {
			case unicode.IsDigit(currChar): // Если первая цифра, то выйти с ошибкой
				//return "", errors.New("wrong string, it can't start from digit")
				return "", ErrInvalidString
			case notLastChar && unicode.IsDigit(rune(str[pos+1])): // Если следующий символ цифра, то обработать на следующем цикле
				break
			case string(currChar) == backSlash: // Если эскейп экран, то обработать на следующем цикле
				break
			default: // Если не цифра, эскейп экран и следующий символ не цифра, то записываем в выходную строку
				strOut.WriteString(string(currChar))
			}

		// Обработка символов кроме первого
		case pos != 0:
			switch {
			case unicode.IsDigit(currChar):
				// Проверка что это не число
				if notLastChar && unicode.IsDigit(rune(str[pos+1])) || unicode.IsDigit(rune(str[pos-1])) {
					//	return "", errors.New("wrong string, it can be only digit, not numeric")
					return "", ErrInvalidString
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
			case string(currChar) == backSlash:
				// Проверка на эскейп последовательность
				if notLastChar {
					if unicode.IsDigit(rune(str[pos+1])) || string(str[pos+1]) == backSlash {
						isDouble = true
						break
					}
				}
				if isDouble {
					break
				}
			default:
				if string(prevChar) == backSlash {
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
