// This package provides the service layer of the application.
// It includes methods for extracting, processing, and validating
// URL parameters and request bodies. Once validated, requests are
// forwarded to the appropriate methods in separate packages.

package service

import (
	"deck-api/pkg/deck"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "deck-api/pkg/responses"
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

// GetNewDeck
// Route for getting a new deck
// @Summary Get a new deck of cards
// @Description Returns a new deck of cards with optional shuffle and cards parameter
// @Accept json
// @Produce json
// @Param shuffle query bool false "Shuffle the deck"
// @Param cards query string false "Specific card codes"
// @Success 200 {object} responses.NewDeckResponse
// @Router /deck [get]
func (s *DeckService) GetNewDeck(context *gin.Context) {
	log.Println("new deck requested")
	var dk *deck.Deck
	var cardCodes []string
	if cardsParam := context.Query("cards"); cardsParam != "" {
		log.Println("specific cards")
		cardCodes = strings.Split(cardsParam, ",")
	}
	if strings.ToLower(context.Query("shuffle")) == "true" {
		log.Println("shuffled deck requested")
		dk = deck.NewDeck(true, cardCodes)
	} else {
		dk = deck.NewDeck(false, cardCodes)
	}
	log.Println("saving deck to decks map")
	s.decks[dk.ID] = dk
	log.Println("deck created successfully")
	context.JSON(http.StatusOK, gin.H{
		"deck_id":   dk.ID,
		"shuffled":  dk.Shuffle,
		"remaining": len(dk.Cards),
	})
}

// OpenDeck
// Route for opening an existing deck
// @Summary Open a deck of cards
// @Description Returns a deck of cards
// @Accept json
// @Produce json
// @Param id path string true "Deck ID"
// @Success 200 {object} responses.OpenDeckResponse
// @Failure 400 {object} responses.ErrorResponse
// @Router /deck/open/{id} [get]
func (s *DeckService) OpenDeck(context *gin.Context) {
	log.Println("open deck requested")
	deckID := context.Param("id")
	if _, ok := s.decks[deckID]; !ok {
		log.Println("deck not found")
		context.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("No deck found with id %s", deckID),
		})
		return
	}
	dk := s.decks[deckID]
	log.Println("deck returned")
	context.JSON(http.StatusOK, dk.Open())
}

// DrawFromDeck
// Route for drawing a card (or multiple cards) from an existing deck
// @Summary Draw cards from the deck
// @Description Returns a deck of cards drawn
// @Accept json
// @Produce json
// @Param id path string true "Deck ID"
// @Param count query string false "Number of cards you want to draw from the deck"
// @Success 200 {object} responses.DrawDeckResponse
// @Failure 400 {object} responses.ErrorResponse
// @Router /deck/draw/{id} [get]
func (s *DeckService) DrawFromDeck(context *gin.Context) {
	log.Println("draw from deck requested")
	deckID := context.Param("id")
	if _, ok := s.decks[deckID]; !ok {
		log.Println("deck not found")
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
			log.Println("incorrect type of count")
			context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid count parameter"})
			return
		}
		log.Println("drawing ", count, " cards from ", deckID)
		cards, cardErr = dk.DrawCard(count)
	} else {
		log.Println("drawing ", 1, " cards from ", deckID)
		cards, cardErr = dk.DrawCard(1)
	}
	if cardErr != nil {
		log.Println("issue with card drawing, error: ", cardErr.Error())
		context.JSON(http.StatusBadRequest, gin.H{
			"message": cardErr.Error(),
		})
		return
	}
	log.Println("draw from deck successful")
	context.JSON(http.StatusOK, gin.H{"cards": cards})
}