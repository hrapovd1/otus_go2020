package main

import (
	"errors"
	"fmt"
	"log"
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
	isDouble := false          //For escaped symbols like \n

	// Loop work with each symbol
	for pos := 0; pos < len(str); pos++ {

		switch {
		// Check first symbol to digital and \
		case pos == 0:
			switch {
			case unicode.IsDigit(rune(str[pos])): // check isn't first digit
				return "", errors.New("wrong string, it can't start from digit")
			case unicode.IsDigit(rune(str[pos+1])): // check if next is digit than need to unpack symbol
				break
			case string(str[pos]) == `\`: // check if it's \ than need to check next symbol
				break
			default: // if it isn't digit, \ or next digit than simple write to out string
				strOut.WriteString(string(str[pos]))
			}

		// Work with another symbols exclude first
		case pos != 0:
			notLastChar := true    // Flag is not last char
			if pos == len(str)-1 { // If it's last char that set flag
				notLastChar = false
			}
			prevChar, currChar := str[pos-1], str[pos]

			switch {
			// if it's digit
			case unicode.IsDigit(rune(currChar)):
				// Check isn't numeric
				if notLastChar && unicode.IsDigit(rune(str[pos+1])) || unicode.IsDigit(rune(str[pos-1])) {
					return "", errors.New("wrong string, it can be only digit, not numeric")
				}
				// convert string digit to int
				currDigit, err := strconv.Atoi(string(currChar))
				if err != nil {
					return "", err
				}

				switch {
				// if digit 0 than skip previous symbol
				case currDigit == 0:
					break
				// if digit more than 0 - unpack it
				case currDigit > 0:
					writedChar := string(prevChar)
					// if previos was escape sequence
					if isDouble {
						writedChar = str[pos-2 : pos]
					}
					for count := currDigit; count > 0; count-- {
						// unpack and write to output
						strOut.WriteString(writedChar)
					}
				}
			// if it's \ than processed it
			case string(currChar) == `\`:
				// if previos was \
				if string(prevChar) == `\` {
					strOut.WriteString(string(prevChar))
				}
				break
			// if it isn't digit, \ or escape
			default:
				if string(prevChar) == `\` {
					// if it's escape sequence set flag
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

func main() {
	const inStr string = `\_\3\:wtest2S\\t` // Input string
	unpackStr, err := Unpack(inStr)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Input stting: %s\nUnpacked string: %s\n", inStr, unpackStr)
}
