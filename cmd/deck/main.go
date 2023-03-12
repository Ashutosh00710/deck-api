// This package contains the main file main.go that initializes
// the major dependencies and starts the server to handle incoming
// requests.

package main

import (
	"deck-api/pkg/server"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	server.Register(engine)
	log.Fatal(engine.Run())
}