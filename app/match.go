package go_regex

func Match(stringToMatch string, exp Expression) MatchTree {
	iter := CreateIterator(stringToMatch)
	return exp(&iter)
}

func MatchIter(iter *Iterator, exp Expression) MatchTree {
	return exp(iter)
}
