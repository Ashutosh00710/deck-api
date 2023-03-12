package deck

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	// Test creating a new deck without shuffling
	deck := NewDeck(false, []string{})
	assert.Equal(t, 52, len(deck.Cards), "Expected deck to have 52 cards")

	// Test creating a new deck with shuffling
	shuffledDeck := NewDeck(true, []string{})
	assert.Equal(t, 52, len(shuffledDeck.Cards), "Expected deck to have 52 cards")
	assert.NotEqual(t, deck.Cards[0].Code, shuffledDeck.Cards[0].Code, "Expected shuffled deck to be different from unshuffled deck")

	// Test creating a new deck with custom cards
	customDeck := NewDeck(false, []string{"AS", "KD"})
	assert.Equal(t, 2, len(customDeck.Cards), "Expected custom deck to have 2 cards")
	assert.Equal(t, "AS", customDeck.Cards[0].Code, "Expected custom deck to have card AS")
	assert.Equal(t, "KD", customDeck.Cards[1].Code, "Expected custom deck to have card KD")
}

func TestShuffleCards(t *testing.T) {
	// Test shuffling a deck
	deck := NewDeck(false, []string{})
	originalCards := make([]Card, len(deck.Cards))
	copy(originalCards, deck.Cards)
	shuffleCards(deck.Cards)
	assert.False(t, reflect.DeepEqual(deck.Cards, originalCards), "Expected shuffled deck to be different from original deck")
}

func TestGenerateCards(t *testing.T) {
	// Test generating a deck of cards
	cards := generateCards()
	assert.Equal(t, 52, len(cards), "Expected 52 cards")
}

func TestBuildCustomDeck(t *testing.T) {
	// Test building a custom deck from a list of card codes
	cards := generateCards()
	customCards := buildCustomDeck([]string{"AS", "KD"}, cards)
	assert.Equal(t, 2, len(customCards), "Expected custom deck to have 2 cards")
	assert.Equal(t, "AS", customCards[0].Code, "Expected custom deck to have card AS")
	assert.Equal(t, "KD", customCards[1].Code, "Expected custom deck to have card KD")
}
