package go_regex

import (
	"strings"
)

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

func (mt *MatchTree) acceptVisitor(visit MatchTreeVisitor) {
	visit(mt)
	for _, child := range mt.Children {
		child.acceptVisitor(visit)
	}
}

func (mt *MatchTree) toString() string {
	sb := strings.Builder{}
	toStringRecursive(mt, &sb, "")
	return sb.String()
}

func toStringRecursive(mt *MatchTree, sb *strings.Builder, levelPadding string) {
	levelPadding = levelPadding + "\t\t"
	sb.WriteString(levelPadding)
	sb.WriteString("|")
	sb.WriteString("\n")

	sb.WriteString(levelPadding)
	sb.WriteString("->[")
	if mt.Label != "" {
		sb.WriteString(mt.Label)
		sb.WriteString(":")
	}
	sb.WriteString(mt.Value)
	sb.WriteString("]")
	sb.WriteString("\n")
	for _, child := range mt.Children {
		toStringRecursive(&child, sb, levelPadding)
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
