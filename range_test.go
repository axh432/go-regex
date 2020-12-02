package gogex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRange(t *testing.T) {
	t.Run("when given a string that matches a range of exactly one return true sequence", func(t *testing.T) {
		sequence := SequenceOfCharacters("a")
		exactlyOne(t, sequence, "SequenceOfCharacters")
	})

	t.Run("when given a string that matches a range of exactly one return true set", func(t *testing.T) {
		set := SetOfCharacters("a")
		exactlyOne(t, set, "SetOfCharacters")
	})

	t.Run("when given a string that matches a range of exactly two return true sequence", func(t *testing.T) {
		sequence := SequenceOfCharacters("a")
		exactlyTwo(t, sequence, "SequenceOfCharacters")
	})

	t.Run("when given a string that matches a range of exactly two return true set", func(t *testing.T) {
		set := SetOfCharacters("a")
		exactlyTwo(t, set, "SetOfCharacters")
	})

	t.Run("when given a string that is within a range return true sequence", func(t *testing.T) {
		sequence := SequenceOfCharacters("a")
		withinRange(t, sequence, "SequenceOfCharacters")
	})

	t.Run("when given a string that is within a range return true set", func(t *testing.T) {
		set := SetOfCharacters("a")
		withinRange(t, set, "SetOfCharacters")
	})

	t.Run("when given a string that is greater than a range return false sequence", func(t *testing.T) {
		sequence := SequenceOfCharacters("a")
		greaterThanRange(t, sequence, "SequenceOfCharacters")
	})

	t.Run("when given a string that is greater than a range return false set", func(t *testing.T) {
		set := SetOfCharacters("a")
		greaterThanRange(t, set, "SetOfCharacters")
	})

	t.Run("when given a string that is less than a range return false sequence", func(t *testing.T) {
		sequence := SequenceOfCharacters("a")
		lessThanRange(t, sequence, "SequenceOfCharacters")
	})

	t.Run("when given a string that is less than a range return false set", func(t *testing.T) {
		set := SetOfCharacters("a")
		lessThanRange(t, set, "SetOfCharacters")
	})

	t.Run("when given a string that is empty and a min of 1 return false sequence", func(t *testing.T) {
		sequence := SequenceOfCharacters("a")
		emptyStringMinOfOne(t, sequence, "SequenceOfCharacters")
	})

	t.Run("when given a string that is empty and a min of 1 return false set", func(t *testing.T) {
		set := SetOfCharacters("a")
		emptyStringMinOfOne(t, set, "SetOfCharacters")
	})

	t.Run("zero minimum sequence", func(t *testing.T) {
		sequence := SequenceOfCharacters("a")
		zeroMinimum(t, sequence, "SequenceOfCharacters")
	})

	t.Run("zero minimum set", func(t *testing.T) {
		set := SetOfCharacters("a")
		zeroMinimum(t, set, "SetOfCharacters")
	})

	t.Run("max is infinity sequence", func(t *testing.T) {
		sequence := SequenceOfCharacters("a")
		maxIsInfinity(t, sequence, "SequenceOfCharacters")
	})

	t.Run("max is infinity set", func(t *testing.T) {
		set := SetOfCharacters("a")
		maxIsInfinity(t, set, "SetOfCharacters")
	})

	t.Run("infinite loop", func(t *testing.T) {
		subExp := Range(SetOfCharacters("a"), 0, -1)

		iter := CreateIterator("xxxxx")

		exp := Range(subExp, 1, -1)

		expected := MatchTree{
			IsValid:   false,
			Value:     "",
			Type:      "Range",
			Children:  []MatchTree(nil),
			DebugLine: "Range:[1:-1], InfiniteLoop:subexpression can capture values of 0 length which will cause Range to loop indefinitely",
		}

		matchResult := MatchIter(&iter, exp)

		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iter.index)

	})
}

func exactlyOne(t *testing.T, subExp Expression, Type string) {
	iter := CreateIterator("a")

	exp := Range(subExp, 1, 1)
	matchResult := MatchIter(&iter, exp)
	expected := MatchTree{
		IsValid:   true,
		Value:     "a",
		Type:      "Range",
		Children:  []MatchTree{{IsValid: true, Type: Type, Value: "a"}},
		DebugLine: "",
	}
	require.Equal(t, expected, matchResult)
	require.Equal(t, 1, iter.index)
}

func exactlyTwo(t *testing.T, subExp Expression, Type string) {
	iter := CreateIterator("aa")

	exp := Range(subExp, 2, 2)
	matchResult := MatchIter(&iter, exp)
	expected := MatchTree{
		IsValid:   true,
		Value:     "aa",
		Type:      "Range",
		Children:  []MatchTree{{IsValid: true, Type: Type, Value: "a"}, {IsValid: true, Type: Type, Value: "a"}},
		DebugLine: "",
	}
	require.Equal(t, expected, matchResult)
	require.Equal(t, 2, iter.index)
}

func withinRange(t *testing.T, subExp Expression, Type string) {
	iter := CreateIterator("aaa")

	exp := Range(subExp, 2, 5)
	matchResult := MatchIter(&iter, exp)
	expected := MatchTree{
		IsValid:   true,
		Value:     "aaa",
		Type:      "Range",
		Children:  []MatchTree{{IsValid: true, Type: Type, Value: "a"}, {IsValid: true, Type: Type, Value: "a"}, {IsValid: true, Type: Type, Value: "a"}},
		DebugLine: "",
	}
	require.Equal(t, expected, matchResult)
	require.Equal(t, 3, iter.index)
}

func greaterThanRange(t *testing.T, subExp Expression, Type string) {
	iter := CreateIterator("aa")

	exp := Range(subExp, 1, 1)
	matchResult := MatchIter(&iter, exp)

	expected := MatchTree{
		IsValid:   false,
		Value:     "",
		Type:      "Range",
		Children:  []MatchTree{{IsValid: true, Type: Type, Value: "a"}, {IsValid: true, Type: Type, Value: "a"}},
		DebugLine: "Range:[1:1], NoMatch:number of subexpressions greater than max",
	}
	require.Equal(t, expected, matchResult)
	require.Equal(t, 0, iter.index)
}

func lessThanRange(t *testing.T, subExp Expression, Type string) {
	iter := CreateIterator("a")

	exp := Range(subExp, 2, 3)
	matchResult := MatchIter(&iter, exp)

	expected := MatchTree{
		IsValid:   false,
		Value:     "",
		Type:      "Range",
		Children:  []MatchTree{{IsValid: true, Type: Type, Value: "a"}},
		DebugLine: "Range:[2:3], NoMatch:number of subexpressions less than min",
	}
	require.Equal(t, expected, matchResult)
	require.Equal(t, 0, iter.index)
}

func emptyStringMinOfOne(t *testing.T, subExp Expression, Type string) {
	iter := CreateIterator("")

	exp := Range(subExp, 1, 2)
	matchResult := MatchIter(&iter, exp)

	expected := MatchTree{
		IsValid:   false,
		Value:     "",
		Type:      "Range",
		Children:  []MatchTree{},
		DebugLine: "Range:[1:2], NoMatch:number of subexpressions less than min",
	}
	require.Equal(t, expected, matchResult)
	require.Equal(t, 0, iter.index)
}

func zeroMinimum(t *testing.T, subExp Expression, Type string) {
	iter := CreateIterator("")

	exp := Range(subExp, 0, 1)
	matchResult := MatchIter(&iter, exp)

	expected := MatchTree{
		IsValid:   true,
		Value:     "",
		Type:      "Range",
		Children:  []MatchTree{},
		DebugLine: "",
	}
	require.Equal(t, expected, matchResult)
	require.Equal(t, 0, iter.index)
}

func maxIsInfinity(t *testing.T, subExp Expression, Type string) {
	iter := CreateIterator("aaaaa")

	exp := Range(subExp, 1, -1)

	expected := MatchTree{
		IsValid:   true,
		Value:     "aaaaa",
		Type:      "Range",
		Children:  []MatchTree{{IsValid: true, Type: Type, Value: "a"}, {IsValid: true, Type: Type, Value: "a"}, {IsValid: true, Type: Type, Value: "a"}, {IsValid: true, Type: Type, Value: "a"}, {IsValid: true, Type: Type, Value: "a"}},
		DebugLine: "",
	}

	matchResult := MatchIter(&iter, exp)

	require.Equal(t, expected, matchResult)
	require.Equal(t, 5, iter.index)

}
