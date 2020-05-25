package go_regex

import "strings"

func SequenceOfCharacters(sequence string) Expression {
	return func(iter *Iterator) MatchTree {

		if sequence == "" {
			return invalidMatchTree("", "SequenceOfCharacters", nil, "SequenceOfCharacters:["+sequence+"], NoMatch:sequence of characters is empty")
		}

		sb := strings.Builder{}
		startingIndex := iter.index

		for _, char := range sequence {
			if !iter.HasNext() {
				iter.Reset(startingIndex)
				return invalidMatchTree(sb.String(), "SequenceOfCharacters", nil, "SequenceOfCharacters:["+sequence+"], NoMatch:reached end of string before finished")
			}

			nextRune := iter.Next()
			if char != nextRune {
				iter.Reset(startingIndex)
				return invalidMatchTree(sb.String(), "SequenceOfCharacters", nil, "SequenceOfCharacters:["+sequence+"], NoMatch: '"+sb.String()+string(nextRune)+"' does not match the sequence")
			}
			sb.WriteRune(char)
		}

		return validMatchTree(sb.String(), "SequenceOfCharacters", nil)
	}
}