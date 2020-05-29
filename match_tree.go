package gogex

type MatchTree struct {
	IsValid   bool
	Value     string
	Type      string
	Label     string
	Children  []MatchTree
	DebugLine string
}

type MatchTreeVisitor func(mt *MatchTree)

type TypeCounter struct {
	setOfCharsCount    int
	setOfNotCharsCount int
	seqOfCharsCount    int
	sequenceCount      int
	setCount           int
	rangeCount         int
	labelCount         int
	stringStartCount   int
	stringEndCount     int
}

func (mt *MatchTree) AcceptVisitor(visit MatchTreeVisitor) {
	visit(mt)
	for _, child := range mt.Children {
		child.AcceptVisitor(visit)
	}
}

func validMatchTree(value string, Type string, children []MatchTree) MatchTree {
	return MatchTree{
		IsValid:   true,
		Value:     value,
		Type:      Type,
		Label:     "",
		Children:  children,
		DebugLine: "",
	}
}

func invalidMatchTree(value, Type string, children []MatchTree, debugLine string) MatchTree {
	return MatchTree{
		IsValid:   false,
		Value:     value,
		Type:      Type,
		Label:     "",
		Children:  children,
		DebugLine: debugLine,
	}
}
