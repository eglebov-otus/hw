package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var lastRune rune
	var needEscape bool
	var result strings.Builder

	for _, r := range input {
		if !needEscape {
			if r == 0x5c { // backslash
				needEscape = true

				continue
			}

			if unicode.IsDigit(r) { // digit (multiplier)
				if lastRune == 0 {
					return "", ErrInvalidString
				}

				multiplier, _ := strconv.Atoi(string(r))
				result.WriteString(strings.Repeat(string(lastRune), multiplier-1))
				lastRune = 0

				continue
			}
		}

		result.WriteRune(r)
		lastRune = r
		needEscape = false
	}

	return result.String(), nil
}
