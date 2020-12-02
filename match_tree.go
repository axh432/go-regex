package gogex

type MatchTree struct {
	IsValid   bool
	Value     string
	Type      string
	Labels    []string
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
		cloneParent.Children = append(cloneParent.Children, cloneWithNoChildren(child))
		cloneChild := &cloneParent.Children[len(cloneParent.Children)-1]
		for _, grandchild := range child.Children {
			pruneToLabelsRecursive(&grandchild, cloneChild)
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
		Labels:    mt.Labels,
		Children:  []MatchTree{},
		DebugLine: mt.DebugLine,
	}
}

func validMatchTree(value string, Type string, children []MatchTree) MatchTree {
	return MatchTree{
		IsValid:   true,
		Value:     value,
		Type:      Type,
		Children:  children,
		DebugLine: "",
	}
}

func invalidMatchTree(value, Type string, children []MatchTree, debugLine string) MatchTree {
	return MatchTree{
		IsValid:   false,
		Value:     value,
		Type:      Type,
		Children:  children,
		DebugLine: debugLine,
	}
}
