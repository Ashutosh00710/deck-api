// This package provides the server layer of the application.
// It includes the HTTP server initialization and routing
// definitions. The package also includes corresponding test files.

package server

import (
	_ "deck-api/docs"
	"deck-api/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func Register(router *gin.Engine) {
	log.Println("START: registering routes")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	registerDeckHandlers(router.Group("/deck"))
	log.Println("DONE: registering routes")
}

func registerDeckHandlers(router *gin.RouterGroup) {
	deckService := service.NewDeckService()
	log.Println("START: registering deck routes")

	router.GET("", deckService.GetNewDeck)
	router.GET("/open/:id", deckService.OpenDeck)
	router.GET("/draw/:id", deckService.DrawFromDeck)
	log.Println("DONE: registering deck routes")
}