package go_regex

import (
	"errors"
	"strings"
	"unicode"
)

func SetOfCharacters(characters string) Expression {
	return func(iter *Iterator) MatchTree {
		if !iter.HasNext() {
			return invalidMatchTree("", "SetOfCharacters", nil, "SetOfCharacters:["+characters+"], NoMatch:reached end of string before finished")
		}

		startingIndex := iter.index
		nextRune := iter.Next()
		for _, char := range characters {
			if char == nextRune {
				return validMatchTree(string(nextRune), "SetOfCharacters", nil)
			}
		}

		iter.Reset(startingIndex)
		return invalidMatchTree("", "SetOfCharacters", nil, "SetOfCharacters:["+characters+"], NoMatch: '"+string(nextRune)+"' not found in set")
	}
}

func SetOfNotCharacters(characters string) Expression {
	return func(iter *Iterator) MatchTree {
		if !iter.HasNext() {
			return invalidMatchTree("", "SetOfNotCharacters", nil, "SetOfNotCharacters:["+characters+"], NoMatch:reached end of string before finished")
		}

		startingIndex := iter.index
		nextRune := iter.Next()
		for _, char := range characters {
			if char == nextRune {
				iter.Reset(startingIndex)
				return invalidMatchTree("", "SetOfNotCharacters", nil, "SetOfNotCharacters:["+characters+"], NoMatch: '"+string(nextRune)+"' found in set")
			}
		}

		return validMatchTree(string(nextRune), "SetOfNotCharacters", nil)
	}
}

func GetSetOfLetters(from rune, to rune) (Expression, error) {
	str, err := GetStringOfLetters(from, to)
	if err != nil {
		return nil, err
	}
	return SetOfCharacters(str), nil
}

func GetSetOfDigits(from rune, to rune) (Expression, error) {
	str, err := GetStringOfDigits(from, to)
	if err != nil {
		return nil, err
	}
	return SetOfCharacters(str), nil
}


func GetStringOfDigits(from rune, to rune) (string, error) {

	if !bothDigits(from, to) {
		return "", errors.New("not all the runes provided were digits")
	}

	return getRunesFromString("0123456789", from, to), nil
}

func GetStringOfLetters(from rune, to rune) (string, error) {
	if !bothLetters(from, to) {
		return "", errors.New("not all the runes provided were letters")
	}
	if bothUpper(from, to) {
		return getRunesFromString("ABCDEFGHIJKLMNOPQRSTUVWXYZ", from, to), nil
	} else if bothLower(from, to) {
		return getRunesFromString("abcdefghijklmnopqrstuvwxyz", from, to), nil
	}
	uppers := getRunesFromString("ABCDEFGHIJKLMNOPQRSTUVWXYZ", unicode.ToUpper(from), unicode.ToUpper(to))
	lowers := getRunesFromString("abcdefghijklmnopqrstuvwxyz", unicode.ToLower(from), unicode.ToLower(to))
	return uppers + lowers, nil
}

func bothDigits(from rune, to rune) bool {
	return unicode.IsDigit(from) && unicode.IsDigit(to)
}

func bothLetters(from rune, to rune) bool {
	return unicode.IsLetter(from) && unicode.IsLetter(to)
}

func bothLower(from rune, to rune) bool {
	return unicode.IsLower(from) && unicode.IsLower(to)
}

func bothUpper(from rune, to rune) bool {
	return unicode.IsUpper(from) && unicode.IsUpper(to)
}

func getRunesFromString(str string, from rune, to rune) string {
	start := strings.IndexRune(str, from)
	end := strings.IndexRune(str, to)
	return str[start : end+1]
}
