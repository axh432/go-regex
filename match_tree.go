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

func (mt *MatchTree) PruneToLabels(labels ...string) MatchTree {
	root := validMatchTree("ROOT", "Label", []MatchTree{})
	pruneToLabelsRecursive(mt, &root, labels)
	return root
}

func (mt *MatchTree) HasLabel(label string) bool {
	for _, mtLabel := range mt.Labels {
		if mtLabel == label {
			return true
		}
	}
	return false
}

func pruneToLabelsRecursive(child *MatchTree, cloneParent *MatchTree, labels []string) {
	if childIsALabel(child) {
		if specificLabelsAreWanted(labels) {
			if childHasAnyOfTheLabels(child, labels) {
				addChildToTree(child, cloneParent, labels)
			} else {
				skipChild(child, cloneParent, labels)
			}
		} else {
			addChildToTree(child, cloneParent, labels)
		}
	} else {
		skipChild(child, cloneParent, labels)
	}
}

func specificLabelsAreWanted(labels []string) bool {
	return len(labels) > 0
}

func childIsALabel(child *MatchTree) bool {
	return child.Type == "Label"
}

func childHasAnyOfTheLabels(child *MatchTree, labels []string) bool {
	for _, label := range labels {
		if child.HasLabel(label) {
			return true
		}
	}
	return false
}

func skipChild(child *MatchTree, cloneParent *MatchTree, labels []string) {
	for _, grandchild := range child.Children {
		pruneToLabelsRecursive(&grandchild, cloneParent, labels)
	}
}

func addChildToTree(child *MatchTree, cloneParent *MatchTree, labels []string) {
	cloneParent.Children = append(cloneParent.Children, cloneWithNoChildren(child))
	cloneChild := &cloneParent.Children[len(cloneParent.Children)-1] //need the pointer
	for _, grandchild := range child.Children {
		pruneToLabelsRecursive(&grandchild, cloneChild, labels)
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
