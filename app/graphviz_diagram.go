package gogex

import (
	"fmt"
	"strconv"
	"strings"
)

func (mt MatchTree) ToGraphVizDiagram() string {
	counter := TypeCounter{}
	definitions := strings.Builder{}
	links := strings.Builder{}

	definitions.WriteString("digraph D {\n\tnode [shape=plaintext fontname=\"Sans serif\" fontsize=\"8\"];\n")

	toGraphVizDiagramRecursive(&mt, "", &counter, &links, &definitions)

	links.WriteString("\n}")

	return fmt.Sprintf("%s\n%s", definitions.String(), links.String())
}

func toGraphVizDiagramRecursive(mt *MatchTree, parentName string, counter *TypeCounter, links *strings.Builder, definitions *strings.Builder) {

	var name string

	switch mt.Type {
	case "SequenceOfCharacters":
		name = fmt.Sprintf("%s_%d", mt.Type, counter.seqOfCharsCount)
		counter.seqOfCharsCount++
		break
	case "SetOfCharacters":
		name = fmt.Sprintf("%s_%d", mt.Type, counter.setOfCharsCount)
		counter.setOfCharsCount++
		break
	case "Set":
		name = fmt.Sprintf("%s_%d", mt.Type, counter.setCount)
		counter.setCount++
		break
	case "Range":
		name = fmt.Sprintf("%s_%d", mt.Type, counter.rangeCount)
		counter.rangeCount++
		break
	case "Sequence":
		name = fmt.Sprintf("%s_%d", mt.Type, counter.sequenceCount)
		counter.sequenceCount++
		break
	case "Label":
		name = fmt.Sprintf("%s_%d", mt.Type, counter.labelCount)
		counter.labelCount++
		break
	case "SetOfNotCharacters":
		name = fmt.Sprintf("%s_%d", mt.Type, counter.setOfNotCharsCount)
		counter.setOfNotCharsCount++
	case "StringStart":
		name = fmt.Sprintf("%s_%d", mt.Type, counter.stringStartCount)
		counter.stringStartCount++
	case "StringEnd":
		name = fmt.Sprintf("%s_%d", mt.Type, counter.stringEndCount)
		counter.stringEndCount++
	}

	if parentName != "" {
		links.WriteString(fmt.Sprintf("\n\t%s\t->%s;", parentName, name))
	}

	classDef := "\n\n" + `%s [ label=<
   <table border="1" cellborder="0" cellspacing="1">
     <tr><td align="left"><b>%s</b></td></tr>
     <tr><td align="left">IsValid: %t</td></tr>
     <tr><td align="left">Value: %s</td></tr>
     <tr><td align="left">Type: %s</td></tr>
     <tr><td align="left">Label: %s</td></tr>
     <tr><td align="left">DebugLine: %s</td></tr>
   </table>>];`

	definitions.WriteString(fmt.Sprintf(classDef, name, name, mt.IsValid, strconv.Quote(mt.Value), mt.Type, mt.Label, mt.DebugLine))

	for _, child := range mt.Children {
		toGraphVizDiagramRecursive(&child, name, counter, links, definitions)
	}
}
