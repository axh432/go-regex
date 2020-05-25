package go_regex

import (
	"sort"
)

func Set(expressions ...Expression) Expression {
	return func(iter *Iterator) MatchTree {
		if len(expressions) == 0 {
			return invalidMatchTree("", "Set", nil, "Set:[], NoMatch:number of subexpressions is zero")
		}

		startingIndex := iter.GetIndex()
		validMatches := []MatchTree{}
		invalidMatches := []MatchTree{}

		for _, exp := range expressions {
			match := exp(iter)
			if match.IsValid {
				validMatches = append(validMatches, match)
			}else{
				invalidMatches = append(invalidMatches, match)
			}
			iter.Reset(startingIndex)
		}

		if(len(validMatches) > 0){
			sort.Slice(validMatches, func(p, q int) bool {
				return len(validMatches[p].Value) > len(validMatches[q].Value) })

			iter.Reset(startingIndex + len(validMatches[0].Value)) //Todo: if len(validMatches[0].Value) == 0 then parent will loop forever.

			return validMatchTree(validMatches[0].Value, "Set", []MatchTree{validMatches[0]})
		}

		iter.Reset(startingIndex)
		return invalidMatchTree("", "Set", invalidMatches, "Set:[], NoMatch:string does not match the given subexpressions")
	}
}

