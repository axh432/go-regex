package gogex

import (
	"fmt"
	"strings"
)

func Range(exp Expression, min int, max int) Expression {
	return func(iter *Iterator) MatchTree {
		startingIndex := iter.GetIndex()
		matches := collectConsecutiveMatches(iter, exp)

		if matches == nil {
			return invalidMatchTree("", "Range", matches, fmt.Sprintf("Range:[%d:%d], InfiniteLoop:subexpression can capture values of 0 length which will cause Range to loop indefinitely", min, max))
		}

		count := len(matches)

		if countAboveMax(count, max) {
			iter.Reset(startingIndex) //always reset the iterator because this might be the child of a set.
			return invalidMatchTree("", "Range", matches, fmt.Sprintf("Range:[%d:%d], NoMatch:number of subexpressions greater than max", min, max))
		}

		if countBelowMin(count, min) {
			iter.Reset(startingIndex) //always reset the iterator because this might be the child of a set.
			return invalidMatchTree("", "Range", matches, fmt.Sprintf("Range:[%d:%d], NoMatch:number of subexpressions less than min", min, max))
		}

		return validMatchTree(createValue(matches), "Range", matches)
	}
}

func createValue(matches []MatchTree) string {
	sb := strings.Builder{}
	for _, child := range matches {
		sb.WriteString(child.Value)
	}
	return sb.String()
}

func collectConsecutiveMatches(iter *Iterator, exp Expression) []MatchTree {
	matches := []MatchTree{}
	for iter.HasNext() {
		startingIndex := iter.index
		match := exp(iter)
		if match.IsValid {
			if len(match.Value) == 0 {
				return nil
			}
			matches = append(matches, match) //I think children need to be pointers
		} else {
			iter.Reset(startingIndex)
			break
		}
	}
	return matches
}

func countAboveMax(count, max int) bool {
	if max < 0 {
		return false //negative numbers classed as infinity
	} else {
		return count > max
	}
}

func countBelowMin(count, min int) bool {
	return count < min
}
