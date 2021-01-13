package gogex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMatchTree_PruneToLabels(t *testing.T) {
	t.Run("check a straight sequence", func(t *testing.T) {
		str := "ababababac"
		a := Label(SetOfCharacters("a"), "A", "Letter")
		b := Label(SetOfCharacters("b"), "B", "Letter")
		c := Label(SetOfCharacters("c"), "C", "Letter")
		exp := Range(Set(a, b, c), 1, -1)
		tree := Match(str, exp)

		allTheAs := tree.PruneToLabels("A")
		allTheBs := tree.PruneToLabels("B")
		allTheLabels := tree.PruneToLabels()

		require.Equal(t, 5, countChildren(allTheAs.Children))
		require.Equal(t, 4, countChildren(allTheBs.Children))
		require.Equal(t, 10, countChildren(allTheLabels.Children))

		require.True(t, checkAllChildren(allTheAs.Children, "A"))
		require.True(t, checkAllChildren(allTheBs.Children, "B"))
		require.True(t, checkAllChildren(allTheLabels.Children, "A", "B", "C", "Letter"))

	})

	t.Run("check a tree sequence", func(t *testing.T) {
		str := "."
		exp := Label(Label(Label(Label(Label(SetOfCharacters("."), "A"), "B"), "A"), "B"), "A")
		tree := Match(str, exp)

		allTheAs := tree.PruneToLabels("A")
		allTheBs := tree.PruneToLabels("B")
		allTheLabels := tree.PruneToLabels()

		require.Equal(t, 3, countChildren(allTheAs.Children))
		require.Equal(t, 2, countChildren(allTheBs.Children))
		require.Equal(t, 5, countChildren(allTheLabels.Children))

		require.True(t, checkAllChildren(allTheAs.Children, "A"))
		require.True(t, checkAllChildren(allTheBs.Children, "B"))
		require.True(t, checkAllChildren(allTheLabels.Children, "A", "B"))
	})

}

func countChildren(children []MatchTree) int {
	totalChildren := len(children)
	for _, child := range children {
		totalChildren += countChildren(child.Children)
	}
	return totalChildren
}

func checkAllChildren(children []MatchTree, labels ...string) (approvedLabel bool) {
	approvedLabel = true
	for _, child := range children {
		approvedLabel = isAnApprovedLabel(child.Labels[0], labels) && checkAllChildren(child.Children, labels...)
	}
	return approvedLabel
}

func isAnApprovedLabel(label string, labels []string) bool {
	for _, l := range labels {
		if l == label {
			return true
		}
	}
	return false
}
