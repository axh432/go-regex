package gogex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSequenceOfCharacters(t *testing.T) {

	t.Run("when given an exact string of characters this expression should return true", func(t *testing.T) {
		iter := CreateIterator("1234567890qwertyuiop[]asdfghjkl;'\\zxcvbnm`,./!@£$%^&*()_+{}|:?><~")
		exp := SequenceOfCharacters("1234567890qwertyuiop[]asdfghjkl;'\\zxcvbnm`,./!@£$%^&*()_+{}|:?><~")
		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid:   true,
			Value:     "1234567890qwertyuiop[]asdfghjkl;'\\zxcvbnm`,./!@£$%^&*()_+{}|:?><~",
			Type:	   "SequenceOfCharacters",
			Label:     "",
			Children:  nil,
			DebugLine: "",
		}

		require.Equal(t, expected, matchResult)
		require.Equal(t, 65, iter.index)
	})

	t.Run("when given a string of characters that differs in the beginning return false", func(t *testing.T) {
		iter := CreateIterator("a")
		exp := SequenceOfCharacters("b")
		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid:   false,
			Value:     "",
			Type:	   "SequenceOfCharacters",
			Label:     "",
			Children:  nil,
			DebugLine: "SequenceOfCharacters:[b], NoMatch: 'a' does not match the sequence",
		}

		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iter.index)
	})

	t.Run("when given a string of characters that differs in the end return false", func(t *testing.T) {
		iter := CreateIterator("ab")
		exp := SequenceOfCharacters("ac")
		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid:   false,
			Value:     "a",
			Type:	   "SequenceOfCharacters",
			Label:     "",
			Children:  nil,
			DebugLine: "SequenceOfCharacters:[ac], NoMatch: 'ab' does not match the sequence",
		}

		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iter.index)
	})

	t.Run("when given a string of characters that differs in the middle return false", func(t *testing.T) {
		iter := CreateIterator("abc")
		exp := SequenceOfCharacters("adc")
		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid:   false,
			Value:     "a",
			Type:	   "SequenceOfCharacters",
			Label:     "",
			Children:  nil,
			DebugLine: "SequenceOfCharacters:[adc], NoMatch: 'ab' does not match the sequence",
		}

		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iter.index)
	})

	t.Run("when given a string of characters that is longer than the sequence return true", func(t *testing.T) {
		iter := CreateIterator("abcdefg")
		exp := SequenceOfCharacters("abc")
		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid:   true,
			Value:     "abc",
			Type:	   "SequenceOfCharacters",
			Label:     "",
			Children:  nil,
			DebugLine: "",
		}

		require.Equal(t, expected, matchResult)
		require.Equal(t, 3, iter.index)
	})

	t.Run("when given a string of characters that is shorter than the sequence return false", func(t *testing.T) {
		iter := CreateIterator("ab")
		exp := SequenceOfCharacters("abc")
		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid:   false,
			Value:     "ab",
			Type:	   "SequenceOfCharacters",
			Label:     "",
			Children:  nil,
			DebugLine: "SequenceOfCharacters:[abc], NoMatch:reached end of string before finished",
		}

		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iter.index)
	})

	t.Run("when given an empty string return false", func(t *testing.T) {
		iter := CreateIterator("")
		exp := SequenceOfCharacters("abc")
		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid:   false,
			Value:     "",
			Type:	   "SequenceOfCharacters",
			Label:     "",
			Children:  nil,
			DebugLine: "SequenceOfCharacters:[abc], NoMatch:reached end of string before finished",
		}

		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iter.index)
	})

	//this is more of a guard against misuse
	t.Run("when given an empty sequence return false", func(t *testing.T) {
		iter := CreateIterator("a")
		exp := SequenceOfCharacters("")
		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid:   false,
			Value:     "",
			Type:	   "SequenceOfCharacters",
			Label:     "",
			Children:  nil,
			DebugLine: "SequenceOfCharacters:[], NoMatch:sequence of characters is empty",
		}

		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iter.index)
	})
}
