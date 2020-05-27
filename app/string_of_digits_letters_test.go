package go_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetStringOfDigits(t *testing.T) {
	t.Run("when given two digits return the range between them", func(t *testing.T) {
		result, err := GetStringOfDigits('3', '7')
		require.NoError(t, err)
		require.Equal(t, "34567", result)
	})
	t.Run("check the full range 0-9", func(t *testing.T) {
		result, err := GetStringOfDigits('0', '9')
		require.NoError(t, err)
		require.Equal(t, "0123456789", result)
	})
	t.Run("when given a rune that is not a digit return an error", func(t *testing.T) {
		result, err := GetStringOfDigits('g', '9')
		require.EqualError(t, err, "not all the runes provided were digits")
		require.Equal(t, "", result)

		result2, err2 := GetStringOfDigits('0', 'f')
		require.EqualError(t, err2, "not all the runes provided were digits")
		require.Equal(t, "", result2)
	})
}

func TestGetStringOfLetters(t *testing.T) {
	t.Run("when given two letters return the range between them", func(t *testing.T) {
		result, err := GetStringOfLetters('l', 'r')
		require.NoError(t, err)
		require.Equal(t, "lmnopqr", result)
	})
	t.Run("when given two capital letters return the range between them", func(t *testing.T) {
		result, err := GetStringOfLetters('L', 'R')
		require.NoError(t, err)
		require.Equal(t, "LMNOPQR", result)
	})
	t.Run("check the full range a-z", func(t *testing.T) {
		result, err := GetStringOfLetters('a', 'z')
		require.NoError(t, err)
		require.Equal(t, "abcdefghijklmnopqrstuvwxyz", result)
	})
	t.Run("check the full capital range A-Z", func(t *testing.T) {
		result, err := GetStringOfLetters('A', 'Z')
		require.NoError(t, err)
		require.Equal(t, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", result)
	})
	t.Run("when given a mix of lowercase and uppercase letters return both ranges", func(t *testing.T) {
		result, err := GetStringOfLetters('l', 'R')
		require.NoError(t, err)
		require.Equal(t, "LMNOPQRlmnopqr", result)

		result2, err2 := GetStringOfLetters('L', 'r')
		require.NoError(t, err2)
		require.Equal(t, "LMNOPQRlmnopqr", result2)
	})
	t.Run("when given a rune that is not a letter return an error", func(t *testing.T) {
		result, err := GetStringOfLetters('g', '9')
		require.EqualError(t, err, "not all the runes provided were letters")
		require.Equal(t, "", result)

		result2, err2 := GetStringOfLetters('0', 'f')
		require.EqualError(t, err2, "not all the runes provided were letters")
		require.Equal(t, "", result2)
	})
}
