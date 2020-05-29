package gogex

import (
	"fmt"
	"strconv"
	"strings"
)

func (mt MatchTree) ToMermaidDiagram() string {
	counter := TypeCounter{}
	definitions := strings.Builder{}
	links := strings.Builder{}

	links.WriteString("classDiagram")

	toMermaidDiagramRecursive(&mt, "", &counter, &links, &definitions)

	return fmt.Sprintf("%s\n%s", links.String(), definitions.String())
}

func toMermaidDiagramRecursive(mt *MatchTree, parentName string, counter *TypeCounter, links *strings.Builder, definitions *strings.Builder) {

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
		links.WriteString(fmt.Sprintf("\n\t%s-->%s", parentName, name))
	}

	classDef := `
class %s {
	IsValid: %t
	Value: %s
	Type: %s
	Label: %s
	DebugLine: %s
}`

	definitions.WriteString(fmt.Sprintf(classDef, name, mt.IsValid, strconv.Quote(replaceMermaidSensitiveCharacters(mt.Value)), mt.Type, mt.Label, replaceMermaidSensitiveCharacters(mt.DebugLine)))

	for _, child := range mt.Children {
		toMermaidDiagramRecursive(&child, name, counter, links, definitions)
	}
}

func replaceMermaidSensitiveCharacters(str string) string {
	smallLeftParenthesis := "﹙"
	smallRightParenthesis := "﹚"
	smallLeftCurlyBrace := "﹛"
	smallRightCurlyBrace := "﹜"
	quotationMark := "“"

	str = strings.ReplaceAll(str, "\"", quotationMark)
	str = strings.ReplaceAll(str, "(", smallLeftParenthesis)
	str = strings.ReplaceAll(str, ")", smallRightParenthesis)
	str = strings.ReplaceAll(str, "{", smallLeftCurlyBrace)
	str = strings.ReplaceAll(str, "}", smallRightCurlyBrace)

	return str
}
