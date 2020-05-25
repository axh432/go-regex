package go_regex

type Iterator struct {
	index int
	end   int
	runes []rune
}

func CreateIterator(str string) Iterator {
	runes := []rune(str)
	return Iterator{index: 0, end: len(runes), runes: runes}
}

func (iter *Iterator) GetIndex() int {
	return iter.index
}

func (iter *Iterator) Reset(newIndex int) {
	iter.index = newIndex
}

func (iter *Iterator) HasPrev() bool {
	return iter.index != 0
}

func (iter *Iterator) HasNext() bool {
	return iter.index != iter.end
}

func (iter *Iterator) Prev() rune {
	iter.index--
	return iter.runes[iter.index]
}

func (iter *Iterator) Next() rune {
	nextRune := iter.runes[iter.index]
	iter.index++
	return nextRune
}
