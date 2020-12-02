package gogex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSequence(t *testing.T) {
	t.Run("when given an exact string of characters this expression should return true", func(t *testing.T) {

		iter := CreateIterator("abc")

		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Sequence(a, b, c)

		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid:   true,
			Value:     "abc",
			Type:      "Sequence",
			Children:  []MatchTree{{IsValid: true, Type: "SetOfCharacters", Value: "a"}, {IsValid: true, Type: "SetOfCharacters", Value: "b"}, {IsValid: true, Type: "SetOfCharacters", Value: "c"}},
			DebugLine: "",
		}

		require.Equal(t, expected, matchResult)
		require.Equal(t, 3, iter.index)
	})

	t.Run("when given a string of characters that differs in the beginning return false", func(t *testing.T) {
		iter := CreateIterator("bbc")

		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Sequence(a, b, c)
		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid:   false,
			Value:     "",
			Type:      "Sequence",
			Children:  []MatchTree{{IsValid: false, Value: "", Type: "SetOfCharacters", DebugLine: "SetOfCharacters:[a], NoMatch: 'b' not found in set"}},
			DebugLine: "Sequence:[], NoMatch:string does not match given subexpression",
		}

		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iter.index)
	})

	t.Run("when given a string of characters that differs in the end return false", func(t *testing.T) {
		iter := CreateIterator("abb")

		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Sequence(a, b, c)
		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid: false,
			Value:   "ab",
			Type:    "Sequence",
			Children: []MatchTree{
				{IsValid: true, Type: "SetOfCharacters", Value: "a"},
				{IsValid: true, Type: "SetOfCharacters", Value: "b"},
				{IsValid: false, Value: "", Type: "SetOfCharacters", DebugLine: "SetOfCharacters:[c], NoMatch: 'b' not found in set"},
			},
			DebugLine: "Sequence:[], NoMatch:string does not match given subexpression",
		}

		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iter.index)
	})

	t.Run("when given a string of characters that differs in the middle return false", func(t *testing.T) {
		iter := CreateIterator("aac")

		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Sequence(a, b, c)
		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid: false,
			Value:   "a",
			Type:    "Sequence",
			Children: []MatchTree{
				{IsValid: true, Type: "SetOfCharacters", Value: "a"},
				{IsValid: false, Value: "", Type: "SetOfCharacters", DebugLine: "SetOfCharacters:[b], NoMatch: 'a' not found in set"},
			},
			DebugLine: "Sequence:[], NoMatch:string does not match given subexpression",
		}

		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iter.index)
	})

	t.Run("when given a string of characters that is longer than the sequence return true", func(t *testing.T) {
		iter := CreateIterator("abcdefg")

		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Sequence(a, b, c)
		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid:   true,
			Value:     "abc",
			Type:      "Sequence",
			Children:  []MatchTree{{IsValid: true, Type: "SetOfCharacters", Value: "a"}, {IsValid: true, Type: "SetOfCharacters", Value: "b"}, {IsValid: true, Type: "SetOfCharacters", Value: "c"}},
			DebugLine: "",
		}

		require.Equal(t, expected, matchResult)
		require.Equal(t, 3, iter.index)
	})

	t.Run("when given a string of characters that is shorter than the sequence return false", func(t *testing.T) {
		iter := CreateIterator("ab")

		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Sequence(a, b, c)
		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid: false,
			Value:   "ab",
			Type:    "Sequence",
			Children: []MatchTree{
				{IsValid: true, Type: "SetOfCharacters", Value: "a"},
				{IsValid: true, Type: "SetOfCharacters", Value: "b"},
				{IsValid: false, Value: "", Type: "SetOfCharacters", DebugLine: "SetOfCharacters:[c], NoMatch:reached end of string before finished"},
			},
			DebugLine: "Sequence:[], NoMatch:string does not match given subexpression",
		}

		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iter.index)
	})

	t.Run("when given an empty string return false", func(t *testing.T) {
		iter := CreateIterator("")

		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Sequence(a, b, c)
		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid: false,
			Value:   "",
			Type:    "Sequence",
			Children: []MatchTree{
				{IsValid: false, Value: "", Type: "SetOfCharacters", DebugLine: "SetOfCharacters:[a], NoMatch:reached end of string before finished"},
			},
			DebugLine: "Sequence:[], NoMatch:string does not match given subexpression",
		}

		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iter.index)
	})

	t.Run("when given an empty sequence return false", func(t *testing.T) {
		iter := CreateIterator("abc")
		exp := Sequence()
		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid:   false,
			Value:     "",
			Type:      "Sequence",
			Children:  nil,
			DebugLine: "Sequence:[], NoMatch:number of subexpressions is zero",
		}

		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iter.index)
	})
}
