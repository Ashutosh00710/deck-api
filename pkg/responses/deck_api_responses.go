// The responses package provides the definitions for the responses
// returned by the various endpoints of the Deck API. It includes
// structs for successful responses, as well as an ErrorResponse
// struct for failed responses.

package responses

import "deck-api/pkg/deck"

// NewDeckResponse represents the response returned by the /deck endpoint
type NewDeckResponse struct {
	// Deck ID
	// example: 1674b1fa-f1d5-4453-b617-88c84b6df7ea
	DeckID string `json:"deck_id"`

	// Number of cards remaining in the deck
	// example: 52
	Remaining int `json:"remaining"`

	// Whether the deck was shuffled or not
	// example: false
	Shuffled bool `json:"shuffled"`
}

// OpenDeckResponse represents the response returned by the /deck/open/:id endpoint
type OpenDeckResponse struct {
	// Deck ID
	// example: 1674b1fa-f1d5-4453-b617-88c84b6df7ea
	DeckID string `json:"deck_id"`

	// Number of cards remaining in the deck
	// example: 52
	Remaining int `json:"remaining"`

	// Whether the deck was shuffled or not
	// example: false
	Shuffled bool `json:"shuffled"`

	// List of cards in the deck
	// example: [{ "rank": "ACE", "suit": "SPADES", "code": "AS" }]
	Cards []deck.Card `json:"cards"`
}

// DrawDeckResponse represents the response returned by the /deck/draw/:id endpoint
type DrawDeckResponse struct {
	// List of cards drawn from the deck
	// example: [{ "rank": "ACE", "suit": "SPADES", "code": "AS" }]
	Cards []deck.Card `json:"cards"`
}

// ErrorResponse represents the error response returned by API endpoint
type ErrorResponse struct {
	// List of cards drawn from the deck
	// example: "something went wrong"
	Message string `json:"message"`
}