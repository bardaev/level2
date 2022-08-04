package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(str string) (string, error) {
	var InvalidString error = errors.New("invalid string")
	if len(str) < 2 {
		return "", InvalidString
	}

	var strRune []rune = []rune(str)
	var lastRune rune
	var result []rune = make([]rune, 0)

	for index, item := range strRune {
		if index == 0 {
			if unicode.IsDigit(item) {
				return "", InvalidString
			}
			lastRune = item
			continue
		}

		if unicode.IsDigit(item) && unicode.IsDigit(lastRune) {
			return "", InvalidString
		}

		if !(unicode.IsLetter(lastRune) && unicode.IsDigit(item)) {
			if unicode.IsLetter(item) && unicode.IsLetter(lastRune) {
				result = append(result, lastRune)
			}
			lastRune = item
			continue
		}

		num, err := strconv.Atoi(string(item))
		if err != nil {
			return "", InvalidString
		}
		var strRepeat string = strings.Repeat(string(lastRune), num)
		var strRepeatRune []rune = []rune(strRepeat)
		result = append(result, strRepeatRune...)
		lastRune = item
	}
	if unicode.IsLetter(lastRune) {
		result = append(result, lastRune)
	}
	return string(result), nil
}
