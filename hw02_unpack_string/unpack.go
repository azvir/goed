package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	ra := []rune(s)

	if len(ra) == 0 {
		return "", nil
	}
	if unicode.IsDigit(ra[0]) {
		return "", ErrInvalidString
	}
	var b strings.Builder
	for i, r := range ra {
		if unicode.IsDigit(r) {
			if unicode.IsDigit(ra[i+1]) {
				return "", ErrInvalidString
			}
			continue
		}
		if i == len(ra)-1 {
			b.WriteRune(r)
			break
		}
		nextRune := ra[i+1]
		if unicode.IsDigit(nextRune) {
			repeat, _ := strconv.Atoi(string(nextRune))
			b.WriteString(strings.Repeat(string(r), repeat))
		} else {
			b.WriteRune(r)
		}
	}
	return b.String(), nil
}
