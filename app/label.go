package go_regex

func Label(exp Expression, label string) Expression {
	return func(iter *Iterator) MatchTree {
		match := exp(iter)
		return MatchTree{
			IsValid:   match.IsValid,
			Value:     match.Value,
			Type: 	   "Label",
			Label:     label,
			Children:  []MatchTree{match},
		}
	}
}
