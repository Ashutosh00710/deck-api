package deck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDrawCard(t *testing.T) {
	// Create a new deck without shuffling
	deck := NewDeck(false, []string{})

	// Draw one card from the deck
	cards, err := deck.DrawCard(1)
	assert.NoError(t, err, "Unexpected error while drawing cards")
	assert.Equal(t, 1, len(cards), "Expected to draw 1 card")
	assert.Equal(t, 51, len(deck.Cards), "Expected deck to have 51 cards")

	// Draw all remaining cards from the deck
	cards, err = deck.DrawCard(51)
	assert.NoError(t, err, "Unexpected error while drawing cards")
	assert.Equal(t, 51, len(cards), "Expected to draw 51 cards")
	assert.Empty(t, deck.Cards, "Expected deck to be empty")

	// Try to draw one more card from the empty deck
	_, err = deck.DrawCard(1)
	assert.Error(t, err, "Expected an error")

	// Try to draw more cards than the deck has
	_, err = deck.DrawCard(52)
	assert.Error(t, err, "Expected an error")
}
