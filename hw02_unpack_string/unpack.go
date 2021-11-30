package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	if len(input) == 0 {
		return "", nil
	}

	runeInput := []rune(input)
	var output strings.Builder

	for i := 0; i < len(runeInput); i++ {
		currentRune := runeInput[i]
		if unicode.IsSpace(currentRune) {
			return "", ErrInvalidString
		}

		if i == 0 && unicode.IsDigit(currentRune) {
			return "", ErrInvalidString
		}

		nextRune, err := getNextRune(runeInput, i)
		if err != nil {
			if unicode.IsDigit(currentRune) {
				continue
			}
		}

		if unicode.IsDigit(currentRune) && unicode.IsDigit(nextRune) {
			return "", ErrInvalidString
		}

		num, zeroError := runeToIntWithoutZero(nextRune)
		if zeroError != nil || unicode.IsDigit(currentRune) {
			continue
		}

		if unicode.IsDigit(nextRune) && num != 0 {
			output.Write([]byte(strings.Repeat(string(currentRune), num)))
			continue
		}

		output.WriteRune(currentRune)
	}

	return output.String(), nil
}

func runeToIntWithoutZero(checkedRune rune) (int, error) {
	if !unicode.IsDigit(checkedRune) {
		return 0, nil
	}

	num, _ := strconv.Atoi(string(checkedRune))

	if num == 0 {
		return 0, errors.New("you can not add the number 0")
	}

	return num, nil
}

func getNextRune(runeInput []rune, i int) (rune, error) {
	if i == len(runeInput)-1 {
		return runeInput[i], errors.New("last element")
	}

	return runeInput[i+1], nil
}
