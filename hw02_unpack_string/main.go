package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(str string) (string, error) {
	// Check empty string
	if len(str) == 0 {
		return "", nil
	}

	var strOut strings.Builder //Output string

	// Loop work with each symbol
	for pos := 0; pos < len(str); pos++ {
		switch {
		case pos == 0:
			switch {
			case unicode.IsDigit(rune(str[pos])): // check isn't first digit
				return "", errors.New("wrong string, it can't start from digit")
			case unicode.IsDigit(rune(str[pos+1])):
				break
			case unicode.IsSymbol('\\'):
				break
			default:
				strOut.WriteString(string(str[pos]))
			}

		case pos != 0:
			switch {
			case unicode.IsDigit(rune(str[pos])):
				// Check isn't numeric
				// TODO: check str[pos] to last symbol in string
				if unicode.IsDigit(rune(str[pos-1])) || unicode.IsDigit(rune(str[pos+1])) {
					return "", errors.New("wrong string, it can be only digit, not numeric")
				}

				prevChar, currChar := str[pos-1], str[pos]
				currDigit, err := strconv.Atoi(string(currChar))
				if err != nil {
					return "", err
				}

				switch {
				case currDigit == 0:
					break
				case currDigit > 0:
					for count := currDigit; count > 0; count-- {
						strOut.WriteString(string(prevChar))
					}
				}
			default:
				strOut.WriteString(string(str[pos]))
			}
			/*
				switch {
				case "0" == strconv.FormatUint(uint64(str[pos]), 10): // if digit is zero than skip symbol
					continue
				}
			*/
		}
	}
	return strOut.String(), nil
}

func main() {
	const inStr string = "test2Str" // Input string
	fmt.Printf("Unpacked string: %s", Unpack(inStr))
}
