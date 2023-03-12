// This package provides the server layer of the application.
// It includes the HTTP server initialization and routing
// definitions. The package also includes corresponding test files.

package server

import (
	"deck-api/pkg/service"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	registerDeckHandlers(router.Group("/deck"))
}

func registerDeckHandlers(router *gin.RouterGroup) {
	deckService := service.NewDeckService()

	router.GET("", deckService.GetNewDeck)
	router.GET("/open/:id", deckService.OpenDeck)
	router.GET("/draw/:id", deckService.DrawFromDeck)
}
