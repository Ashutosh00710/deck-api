// This package provides the service layer of the application.
// It includes methods for extracting, processing, and validating
// URL parameters and request bodies. Once validated, requests are
// forwarded to the appropriate methods in separate packages.

package service

import (
	"github.com/gin-gonic/gin"
)

type DeckService struct {}

func NewDeckService() *DeckService {
	return &DeckService{}
}

func (s *DeckService) GetNewDeck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message":   "success",
	})
}