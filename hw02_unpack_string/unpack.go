package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	if len(input) == 0 {
		return "", nil
	}

	var (
		output   strings.Builder
		prevRune rune
	)

	for i, r := range input {
		validateErr := validate(i, prevRune, r)
		if validateErr != nil {
			return "", validateErr
		}

		if unicode.IsDigit(r) {
			num, _ := strconv.Atoi(string(r))
			if num == 0 {
				outputString := output.String()
				_, size := utf8.DecodeLastRuneInString(outputString)
				output.Reset()
				output.WriteString(outputString[:len(outputString)-size])
				continue
			}

			output.Write([]byte(strings.Repeat(string(prevRune), num-1)))
			prevRune = r
			continue
		}

		prevRune = r

		if !unicode.IsDigit(r) {
			output.WriteRune(r)
		}
	}

	return output.String(), nil
}

func validate(index int, prevR, r rune) error {
	conditions := []bool{
		unicode.IsSpace(r),
		index == 0 && unicode.IsDigit(r),
		unicode.IsDigit(prevR) && unicode.IsDigit(r),
	}

	for _, condition := range conditions {
		if condition {
			return ErrInvalidString
		}
	}

	return nil
}
