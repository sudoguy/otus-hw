package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func hasManyDigits(text string) bool {
	r := regexp.MustCompile(`\d{2,}`)

	digits := r.FindString(text)

	return digits != ""
}

func Unpack(text string) (string, error) {
	var builder strings.Builder
	var prevSymbol string

	firstSymbol, _ := utf8.DecodeRuneInString(text)
	// Если несколько цифр подряд или строка начинается с цифры, то возвращаем ошибку
	if hasManyDigits(text) || unicode.IsDigit(firstSymbol) {
		return "", ErrInvalidString
	}

	for _, r := range text {
		currentSymbol := string(r)

		if unicode.IsDigit(r) {
			// если цифра, то пишем предыдущую букву count-1 раз, т.к. 1 раз букву мы уже записали
			count, _ := strconv.Atoi(currentSymbol)
			builder.WriteString(strings.Repeat(prevSymbol, count-1))
		} else {
			// букву пишем всегда
			builder.WriteString(currentSymbol)
		}
		prevSymbol = currentSymbol
	}
	return builder.String(), nil
}
