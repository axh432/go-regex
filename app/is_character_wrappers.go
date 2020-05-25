package go_regex

import (
	"unicode"
)

type IsCharacterFunction func(r rune) bool

var (
	Whitespace  = createSetFromIsCharacterFunction(unicode.IsSpace, "Whitespace")
	Number      = createSetFromIsCharacterFunction(unicode.IsNumber, "Number")
	Letter      = createSetFromIsCharacterFunction(unicode.IsLetter, "Letter")
	Punctuation = createSetFromIsCharacterFunction(unicode.IsPunct, "Punctuation")
	Symbol      = createSetFromIsCharacterFunction(unicode.IsSymbol, "Symbol")
)

func createSetFromIsCharacterFunction(isCharacterFunction IsCharacterFunction, charSetName string) Expression {
	return func(iter *Iterator) MatchTree {

		startingIndex := iter.index

		if !iter.HasNext() {
			return invalidMatchTree("", "SetOfCharacters", nil, "SetOfCharacters:["+charSetName+"], NoMatch:reached end of string before finished")
		}

		nextRune := iter.Next()
		if isCharacterFunction(nextRune) {
			return validMatchTree(string(nextRune), "SetOfCharacters", nil)
		}

		iter.Reset(startingIndex)
		return invalidMatchTree("", "SetOfCharacters", nil, "SetOfCharacters:["+charSetName+"], NoMatch: '"+string(nextRune)+"' not found in set")
	}
}
