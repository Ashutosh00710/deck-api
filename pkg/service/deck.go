// This package provides the service layer of the application.
// It includes methods for extracting, processing, and validating
// URL parameters and request bodies. Once validated, requests are
// forwarded to the appropriate methods in separate packages.

package service

import (
	"deck-api/pkg/deck"
	"fmt"
	"net/http"
	"strconv"
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

func (s *DeckService) OpenDeck(context *gin.Context) {
	deckID := context.Param("id")
	if _, ok := s.decks[deckID]; !ok {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("No deck found with id %s", deckID),
		})
		return
	}
	dk := s.decks[deckID]
	context.JSON(http.StatusOK, dk.Open())
}

func (s *DeckService) DrawFromDeck(context *gin.Context) {
	deckID := context.Param("id")
	if _, ok := s.decks[deckID]; !ok {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("No deck found with id %s", deckID),
		})
		return
	}
	dk := s.decks[deckID]
	var cards []deck.Card
	var cardErr error
	if countStr := context.Query("count"); countStr != "" {
		count, err := strconv.Atoi(countStr)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid count parameter"})
			return
		}
		cards, cardErr = dk.DrawCard(count)
	} else {
		cards, cardErr = dk.DrawCard(1)
	}
	if cardErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": cardErr.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{"cards": cards})
}
