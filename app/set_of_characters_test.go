package go_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSetOfCharacters(t *testing.T) {
	t.Run("when given a character that is in the set return true", func(t *testing.T) {
		exp := SetOfCharacters("abc")

		iterA := CreateIterator("a")
		matchResultA := MatchIter(&iterA, exp)
		expectedA := MatchTree{Value: "a", DebugLine: "", Type: "SetOfCharacters", IsValid: true}
		require.Equal(t, matchResultA, expectedA)
		require.Equal(t, 1, iterA.index)

		iterB := CreateIterator("b")
		matchResultB := MatchIter(&iterB, exp)
		expectedB := MatchTree{Value: "b", DebugLine: "", Type: "SetOfCharacters", IsValid: true}
		require.Equal(t, matchResultB, expectedB)
		require.Equal(t, 1, iterB.index)

		iterC := CreateIterator("c")
		matchResultC := MatchIter(&iterC, exp)
		expectedC := MatchTree{Value: "c", DebugLine: "", Type: "SetOfCharacters", IsValid: true}
		require.Equal(t, matchResultC, expectedC)
		require.Equal(t, 1, iterC.index)
	})

	t.Run("when given a character that is not in the set return false", func(t *testing.T) {
		exp := SetOfCharacters("abc")
		iterD := CreateIterator("d")
		matchResultD := MatchIter(&iterD, exp)
		expectedD := MatchTree{
			Value:     "",
			Type: "SetOfCharacters",
			DebugLine: "SetOfCharacters:[abc], NoMatch: 'd' not found in set",
			IsValid:   false,
		}
		require.Equal(t, matchResultD, expectedD)
		require.Equal(t, 0, iterD.index)
	})

	t.Run("when given a string and the first character matches return true", func(t *testing.T) {
		exp := SetOfCharacters("abc")

		iterA := CreateIterator("athguy")
		matchResult := MatchIter(&iterA, exp)
		expected := MatchTree{
			IsValid:   true,
			Type: "SetOfCharacters",
			Value:     "a",
			DebugLine: "",
		}
		require.Equal(t, expected, matchResult)
		require.Equal(t, 1, iterA.index)
	})

	t.Run("when given a string and the first character does not match return false", func(t *testing.T) {
		exp := SetOfCharacters("abc")

		iterX := CreateIterator("xthguy")
		matchResult := MatchIter(&iterX, exp)
		expected := MatchTree{
			IsValid:   false,
			Type: "SetOfCharacters",
			Value:     "",
			DebugLine: "SetOfCharacters:[abc], NoMatch: 'x' not found in set",
		}
		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iterX.index)
	})

	t.Run("when given an empty string return false", func(t *testing.T) {
		exp := SetOfCharacters("abc")

		iter := CreateIterator("")
		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid:   false,
			Type: "SetOfCharacters",
			Value:     "",
			DebugLine: "SetOfCharacters:[abc], NoMatch:reached end of string before finished",
		}
		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iter.index)
	})

	t.Run("when given an empty set return false", func(t *testing.T) {
		exp := SetOfCharacters("")

		iter := CreateIterator("a")
		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid:   false,
			Type: "SetOfCharacters",
			Value:     "",
			DebugLine: "SetOfCharacters:[], NoMatch: 'a' not found in set",
		}
		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iter.index)
	})
}