// This package provides the service layer of the application.
// It includes methods for extracting, processing, and validating
// URL parameters and request bodies. Once validated, requests are
// forwarded to the appropriate methods in separate packages.

package service

import (
	"deck-api/pkg/deck"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type DeckService struct {
	decks map[string]*deck.Deck
}

func NewDeckService() *DeckService {
	return &DeckService{
		decks: map[string]*deck.Deck{},
	}
}

func (s *DeckService) GetNewDeck(context *gin.Context) {
	var dk *deck.Deck
	var cardCodes []string
	if cardsParam := context.Query("cards"); cardsParam != "" {
		cardCodes = strings.Split(cardsParam, ",")
	}
	if strings.ToLower(context.Query("shuffle")) == "true" {
		dk = deck.NewDeck(true, cardCodes)
	} else {
		dk = deck.NewDeck(false, cardCodes)
	}
	s.decks[dk.ID] = dk
	context.JSON(http.StatusOK, gin.H{
		"deck_id":   dk.ID,
		"shuffled":  dk.Shuffle,
		"remaining": len(dk.Cards),
	})
}
