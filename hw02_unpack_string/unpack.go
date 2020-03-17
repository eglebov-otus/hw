package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var prev rune
	var res strings.Builder

	for _, cur := range input {
		if unicode.IsLetter(cur) {
			res.WriteRune(cur)
		}

		if unicode.IsDigit(cur) {
			if unicode.IsDigit(prev) || prev == 0 {
				return "", ErrInvalidString
			}

			multiplier, _ := strconv.Atoi(string(cur))
			multiplier--
			res.WriteString(strings.Repeat(string(prev), multiplier))
		}

		prev = cur
	}

	return res.String(), nil
}
