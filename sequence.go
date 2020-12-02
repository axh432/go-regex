package gogex

import (
	"fmt"
	"strings"
)

func Sequence(expressions ...Expression) Expression {
	return func(iter *Iterator) MatchTree {
		if len(expressions) == 0 {
			return invalidMatchTree("", "Sequence", nil, "Sequence:[], NoMatch:number of subexpressions is zero")
		}

		sb := strings.Builder{}
		matches := []MatchTree{}

		startingIndex := iter.index

		for _, exp := range expressions {
			match := exp(iter)
			matches = append(matches, match)
			if match.IsValid {
				sb.WriteString(match.Value)
			} else {
				iter.Reset(startingIndex)
				debugLine := "Sequence:[], NoMatch:string does not match given subexpression"
				if len(match.Labels) > 0 {
					debugLine = fmt.Sprintf("Sequence:[], NoMatch:string does not match given subexpression: %s", formatLabels(match.Labels))
				}
				return invalidMatchTree(sb.String(), "Sequence", matches, debugLine)
			}
		}

		return validMatchTree(sb.String(), "Sequence", matches)
	}
}
