package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrStringStartWithDigit = errors.New("string starts with a digit")
	ErrStringEndWithEscape  = errors.New("string ends with escape character")
	ErrInternal             = errors.New("internal error")
)

func UnpackString(s string) (string, error) {
	runes := []rune(s)

	if len(runes) == 0 {
		return "", nil
	}

	if unicode.IsDigit(runes[0]) {
		return "", ErrStringStartWithDigit
	}

	var result strings.Builder

	for i := 0; i < len(runes); i++ {
		char := runes[i]

		if char == '\\' {
			if i+1 < len(runes) {
				i++
				result.WriteRune(runes[i])
				continue
			} else {
				return "", ErrStringEndWithEscape
			}
		}

		if unicode.IsDigit(char) {
			lastChar := runes[i-1]

			var multiplier strings.Builder
			multiplier.WriteRune(char)

			for j := i + 1; j < len(runes); j++ {
				if unicode.IsDigit(runes[j]) {
					multiplier.WriteRune(runes[j])
					i++
				} else {
					break
				}
			}

			strMultiplier := multiplier.String()
			intMultiplier, err := strconv.Atoi(strMultiplier)
			if err != nil {
				return "", ErrInternal
			}

			for k := 0; k < intMultiplier-1; k++ {
				result.WriteRune(lastChar)
			}
			continue

		}

		result.WriteRune(char)
	}

	return result.String(), nil
}

func main() {
	str := `qwe\45`
	res, err := UnpackString(str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(res)
}
