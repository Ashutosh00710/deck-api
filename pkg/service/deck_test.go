package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetNewDeck(t *testing.T) {
	// create a new gin context and deck service
	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	service := NewDeckService()

	service.GetNewDeck(context)
	assert.Len(t, service.decks, 1)

	// add query parameter to context
	context.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
	context.Request.URL.RawQuery = "shuffle=true&cards=AS,2H,3D"
	service.GetNewDeck(context)
	assert.Len(t, service.decks, 2)
}

func TestOpenDeck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	service := NewDeckService()

	service.GetNewDeck(context)
	assert.Len(t, service.decks, 1)

	// test with correct id
	for _, deck := range service.decks {
		context.Request, _ = http.NewRequest(http.MethodGet, "/"+deck.ID, nil)
		service.OpenDeck(context)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	}

	// test with incorrect id
	context, _ = gin.CreateTestContext(httptest.NewRecorder())
	context.Request, _ = http.NewRequest(http.MethodGet, "/123", nil)
	service.OpenDeck(context)
	assert.Equal(t, http.StatusBadRequest, context.Writer.Status())
}

func TestDrawDeck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	service := NewDeckService()

	service.GetNewDeck(context)
	assert.Len(t, service.decks, 1)

	// test with correct id
	for _, deck := range service.decks {
		// with correct count
		context.Request, _ = http.NewRequest(http.MethodGet, "/"+deck.ID+"?count=1", nil)
		service.DrawFromDeck(context)
		assert.Equal(t, http.StatusOK, context.Writer.Status())

		// with incorrect count
		context, _ = gin.CreateTestContext(httptest.NewRecorder())
		context.Request, _ = http.NewRequest(http.MethodGet, "/"+deck.ID+"?count=52", nil)
		service.DrawFromDeck(context)
		assert.Equal(t, http.StatusBadRequest, context.Writer.Status())
	}

	// test with incorrect id
	context, _ = gin.CreateTestContext(httptest.NewRecorder())
	context.Request, _ = http.NewRequest(http.MethodGet, "/123", nil)
	service.OpenDeck(context)
	assert.Equal(t, http.StatusBadRequest, context.Writer.Status())
}