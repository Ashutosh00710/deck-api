// This package provides the core functionality for working with a
// deck of cards. It includes methods for creating a deck, shuffling
// a deck, and drawing cards from a deck. The package also includes
// corresponding test files.

package deck

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Card struct {
	Rank Rank   `json:"rank"`
	Suit Suit   `json:"suit"`
	Code string `json:"code"`
}

type Deck struct {
	ID      string `json:"deck_id"`
	Cards   []Card `json:"cards"`
	Shuffle bool   `json:"shuffle"`
}

// cache to avoid creation of card process
var allCards []Card

// NewDeck creates a new deck with the given parameters.
func NewDeck(shuffle bool, cardCodes []string) *Deck {
	if len(allCards) == 0 {
		allCards = generateCards()
	}
	var cards []Card
	cards = append(cards, allCards...)

	if len(cardCodes) > 0 {
		cards = buildCustomDeck(cardCodes, cards)
	}

	if shuffle {
		shuffleCards(cards)
	}

	return &Deck{
		ID:      uuid.New().String(),
		Cards:   cards,
		Shuffle: shuffle,
	}
}

// shuffleCards shuffles the given cards slice.
func shuffleCards(cards []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}

// generateCards generates all possible cards with their ranks, suits, and codes.
func generateCards() []Card {
	var cards []Card
	ranks := []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
	suits := []Suit{Spades, Diamonds, Clubs, Hearts}
	for _, suit := range suits {
		for _, rank := range ranks {
			card := Card{Rank: rank, Suit: suit, Code: string(rank) + string(suit)[0:1]}
			cards = append(cards, card)
		}
	}
	return cards
}

// buildCustomDeck builds a custom deck using the given card codes and the provided cards slice.
func buildCustomDeck(cardCodes []string, cards []Card) []Card {
	var filteredCards []Card
	cardCodesMap := map[string]bool{}
	for _, cardStr := range cardCodes {
		cardCodesMap[cardStr] = true
	}
	for _, card := range cards {
		if _, ok := cardCodesMap[card.Code]; ok {
			filteredCards = append(filteredCards, card)
		}
	}
	return filteredCards
}

// Rank represents the rank of a playing card.
type Rank string

// Suit represents the suit of a playing card.
type Suit string

// Constants for each possible rank and suit.
const (
	Ace   Rank = "A"
	Two   Rank = "2"
	Three Rank = "3"
	Four  Rank = "4"
	Five  Rank = "5"
	Six   Rank = "6"
	Seven Rank = "7"
	Eight Rank = "8"
	Nine  Rank = "9"
	Ten   Rank = "10"
	Jack  Rank = "J"
	Queen Rank = "Q"
	King  Rank = "K"

	Spades   Suit = "Spades"
	Diamonds Suit = "Diamonds"
	Clubs    Suit = "Clubs"
	Hearts   Suit = "Hearts"
)
