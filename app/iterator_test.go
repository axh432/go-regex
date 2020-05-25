package go_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIterator_Reset(t *testing.T) {
	iter := CreateIterator("F (l r i){+}")
	iter.Reset(3)
	require.Equal(t, "l", string(iter.Next()))
	iter.Reset(0)
	require.Equal(t, "F", string(iter.Next()))
}

func TestIterator_HasPrev(t *testing.T) {
	iter := CreateIterator("F (l r i){+}")
	iter.Next()
	require.True(t, iter.HasPrev())

	iter.Reset(0)
	require.False(t, iter.HasPrev())
}

func TestIterator_HasNext(t *testing.T) {
	iter := CreateIterator("{+}")
	iter.Next()
	require.True(t, iter.HasNext())

	iter.Reset(3)
	require.False(t, iter.HasNext())
}

func TestIterator_Prev(t *testing.T) {
	iter := CreateIterator("F (l r i){+}")
	iter.Reset(3)
	require.Equal(t, "(", string(iter.Prev()))
}

func TestIterator_Next(t *testing.T) {
	iter := CreateIterator("F (l r i){+}")
	require.Equal(t, "F", string(iter.Next()))
	require.Equal(t, " ", string(iter.Next()))
	require.Equal(t, "(", string(iter.Next()))
}
