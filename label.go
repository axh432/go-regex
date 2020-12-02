package gogex

func Label(exp Expression, label ...string) Expression {
	return func(iter *Iterator) MatchTree {
		match := exp(iter)
		return MatchTree{
			IsValid:  match.IsValid,
			Value:    match.Value,
			Type:     "Label",
			Labels:   label,
			Children: []MatchTree{match},
		}
	}
}
