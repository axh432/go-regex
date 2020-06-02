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

func (mt *MatchTree) PruneToLabels() MatchTree {
	clone := cloneWithNoChildren(mt)
	for _, child := range mt.Children {
		pruneToLabelsRecursive(&child, &clone)
	}
	return clone
}

func pruneToLabelsRecursive(child *MatchTree, cloneParent *MatchTree) {
	if child.Type == "Label" {
		childClone := cloneWithNoChildren(child)
		cloneParent.Children = append(cloneParent.Children, childClone)
		for _, grandchild := range child.Children {
			pruneToLabelsRecursive(&grandchild, &childClone)
		}
	} else {
		for _, grandchild := range child.Children {
			pruneToLabelsRecursive(&grandchild, cloneParent)
		}
	}
}

func cloneWithNoChildren(mt *MatchTree) MatchTree {
	return MatchTree{
		IsValid:   mt.IsValid,
		Value:     mt.Value,
		Type:      mt.Type,
		Label:     mt.Label,
		Children:  []MatchTree{},
		DebugLine: mt.DebugLine,
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