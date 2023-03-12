package deck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenShuffledDeck(t *testing.T) {
	// create a shuffled deck
	deck := NewDeck(true, nil)

	// call Open function
	deckData := deck.Open()

	// check if the deck ID is correct
	assert.Equal(t, deck.ID, deckData["deck_id"], "Open function returned incorrect deck ID")

	// check if the shuffle status is correct
	assert.True(t, deckData["shuffled"].(bool), "Open function returned incorrect shuffle status")

	// check if the remaining cards count is correct
	assert.Equal(t, len(deck.Cards), deckData["remaining"].(int), "Open function returned incorrect remaining cards count")

	// check if the cards in the response match the cards in the deck
	assert.Equal(t, deck.Cards, deckData["cards"].([]Card), "Open function returned incorrect cards")
}

func TestOpenUnshuffledDeck(t *testing.T) {
	// create an un-shuffled deck
	deck := NewDeck(false, nil)

	// call Open function
	deckData := deck.Open()

	// check if the deck ID is correct
	assert.Equal(t, deck.ID, deckData["deck_id"], "Open function returned incorrect deck ID")

	// check if the shuffle status is correct
	assert.False(t, deckData["shuffled"].(bool), "Open function returned incorrect shuffle status")

	// check if the remaining cards count is correct
	assert.Equal(t, len(deck.Cards), deckData["remaining"].(int), "Open function returned incorrect remaining cards count")

	// check if the cards in the response match the cards in the deck
	assert.Equal(t, deck.Cards, deckData["cards"].([]Card), "Open function returned incorrect cards")
}
