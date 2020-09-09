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
	const backSlash = '\\'     // Оптимизаци линтера goconst
	strRune := []rune(str)

	// Основной цикл перебора символов
	for pos := 0; pos < len(strRune); pos++ {
		// Для удобства понимания и работы с кодом определяю переменные предыдущего и текущего символов
		var prevChar, currChar rune
		if pos > 0 {
			prevChar, currChar = strRune[pos-1], strRune[pos]
		} else {
			prevChar, currChar = 0, strRune[pos]
		}
		notLastChar := true // Установка флага, что это не последний символ
		// Проверка на последний символ
		if pos == len(strRune)-1 {
			notLastChar = false
		}

		switch {
		// Проверка первого символа на цифру и эскейп экран
		case pos == 0:
			switch {
			case unicode.IsDigit(currChar): // Если первая цифра, то выйти с ошибкой
				return "", ErrInvalidString
			case notLastChar && unicode.IsDigit(strRune[pos+1]):
				break
			case notLastChar && currChar == backSlash:
				isDouble = true
			default: // Если не цифра, эскейп экран и следующий символ не цифра, то записываем в выходную строку
				strOut.WriteRune(currChar)
			}

		// Обработка символов кроме первого
		case pos != 0:
			switch {
			case unicode.IsDigit(currChar):
				// Преобразование цифры из строки в целое
				currDigit, err := strconv.Atoi(string(currChar))
				if err != nil {
					return "", err
				}

				if isDouble {
					switch {
					case unicode.IsDigit(prevChar):
						if currDigit > 0 {
							for count := currDigit; count > 0; count-- {
								strOut.WriteRune(prevChar)
							}
						}
						isDouble = false // После обработки экранирования предыдущей цифры, снимаю флаг
					case prevChar == backSlash:
						switch {
						case notLastChar && unicode.IsDigit(strRune[pos+1]):
							if pos-2 >= 0 { // Проверяю что есть два символа перед текущим
								if strRune[pos-2] == backSlash { // Проверяю что ранее было 2 обратных слеша
									return "", ErrInvalidString // если это так, то это ошибка \\dd
								}
							}
							isDouble = false
						case notLastChar || pos-2 >= 0:
							if pos-2 >= 0 && strRune[pos-2] == backSlash {
								if currDigit > 0 {
									for count := currDigit; count > 0; count-- {
										strOut.WriteRune(prevChar)
									}
								}
							} else {
								strOut.WriteRune(currChar)
							}
							isDouble = false
						}
					default:
						if currDigit > 0 {
							for count := currDigit; count > 0; count-- {
								strOut.WriteString(string(strRune[pos-2 : pos]))
							}
						}
						isDouble = false
					}
				} else {
					// Проверка что это не число
					if notLastChar && unicode.IsDigit(strRune[pos+1]) || unicode.IsDigit(strRune[pos-1]) {
						return "", ErrInvalidString
					} else if currDigit > 0 {
						for count := currDigit; count > 0; count-- {
							strOut.WriteRune(prevChar)
						}
					}
				}
			case currChar == backSlash:
				// Проверка на эскейп последовательность
				if notLastChar {
					if unicode.IsDigit(strRune[pos+1]) || strRune[pos+1] == backSlash {
						isDouble = true
						break
					}
				}
				if isDouble {
					break
				}
			default:
				if prevChar == backSlash {
					// Проверка на эскейп последовательность
					if notLastChar && unicode.IsDigit(strRune[pos+1]) {
						isDouble = true
						break
					}
					strOut.WriteRune(prevChar)
				}
				if notLastChar && unicode.IsDigit(strRune[pos+1]) {
					break
				}
				strOut.WriteRune(currChar)
			}
		}
	}
	return strOut.String(), nil
}
