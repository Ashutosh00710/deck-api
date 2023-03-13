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

var allCards []Card

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

func shuffleCards(cards []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}

func generateCards() []Card {
	var cards []Card
	ranks := []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
	suits := []Suit{Spades, Diamonds, Clubs, Hearts}
	for _, suit := range suits {
		for _, rank := range ranks {
			card := Card{Rank: rank, Suit: suit, Code: string(rank)[0:1] + string(suit)[0:1]}
			cards = append(cards, card)
		}
	}
	return cards
}

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

type Rank string
type Suit string

const (
	Ace   Rank = "ACE"
	Two   Rank = "2"
	Three Rank = "3"
	Four  Rank = "4"
	Five  Rank = "5"
	Six   Rank = "6"
	Seven Rank = "7"
	Eight Rank = "8"
	Nine  Rank = "9"
	Ten   Rank = "10"
	Jack  Rank = "JOKER"
	Queen Rank = "QUEEN"
	King  Rank = "KING"

	Spades   Suit = "SPADES"
	Diamonds Suit = "DIAMONDS"
	Clubs    Suit = "CLUBS"
	Hearts   Suit = "HEARTS"
)
