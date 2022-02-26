package random

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShuffle(t *testing.T) {
	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var deckBeforeShuffle = make([]int, len(deck))

	copy(deckBeforeShuffle, deck)

	err := Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	found := false
	for i := 0; i < len(deck); i++ {
		if deck[i] != deckBeforeShuffle[i] {
			found = true
		}
	}
	require.True(t, found)

	require.NoError(t, err)
}
