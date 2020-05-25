package go_regex

func StringStart(iter *Iterator) MatchTree {
	if iter.index == 0 {
		return validMatchTree("", "StringStart", nil)
	}else{
		return invalidMatchTree("", "StringStart", nil, "StringStart, NoMatch:this is not the start of the string")
	}
}

func StringEnd(iter *Iterator) MatchTree {
	if iter.HasNext() {
		return invalidMatchTree("", "StringEnd", nil, "StringEnd, NoMatch:this is not the end of the string")
	}else{
		return validMatchTree("", "StringEnd", nil)
	}
}
