package server

import (
	"deck-api/pkg/deck"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleDeck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	registerDeckHandlers(router.Group("/deck"))

	// expected response object
	var res struct {
		DeckId    string `json:"deck_id"`
		Shuffled  bool   `json:"shuffled"`
		Remaining int    `json:"remaining"`
	}

	// Test without query parameters
	req, _ := http.NewRequest("GET", "/deck", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	err := json.NewDecoder(resp.Body).Decode(&res)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Nil(t, err)
	assert.Equal(t, 52, res.Remaining)

	// Test with shuffle query parameter
	req, _ = http.NewRequest("GET", "/deck?shuffle=true", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	err = json.NewDecoder(resp.Body).Decode(&res)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Nil(t, err)
	assert.Equal(t, 52, res.Remaining)
	assert.Equal(t, true, res.Shuffled)

	// Test with cards query parameter, with correct card codes
	req, _ = http.NewRequest("GET", "/deck?cards=AS,2H,3C", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	err = json.NewDecoder(resp.Body).Decode(&res)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Nil(t, err)
	assert.Equal(t, 3, res.Remaining)

	// Test with cards query parameter, with one incorrect card code
	req, _ = http.NewRequest("GET", "/deck?cards=AS,2H,4F", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	err = json.NewDecoder(resp.Body).Decode(&res)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Nil(t, err)
	assert.Equal(t, 2, res.Remaining)
}

func TestHandleOpenDeck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	registerDeckHandlers(router.Group("/deck"))

	var res struct {
		DeckId    string `json:"deck_id"`
		Shuffled  bool   `json:"shuffled"`
		Remaining int    `json:"remaining"`
	}

	// Test without query parameters
	req, _ := http.NewRequest("GET", "/deck", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	err := json.NewDecoder(resp.Body).Decode(&res)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Nil(t, err)

	// Test with valid deck ID
	req, _ = http.NewRequest("GET", "/deck/open/"+res.DeckId, nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "deck_id")
	assert.Contains(t, resp.Body.String(), "shuffled")
	assert.Contains(t, resp.Body.String(), "remaining")
	assert.Contains(t, resp.Body.String(), "cards")

	// Test with invalid deck ID
	req, _ = http.NewRequest("GET", "/deck/open/456", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "No deck found")
}

func TestHandleDrawDeck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	registerDeckHandlers(router.Group("/deck"))

	var resDeck struct {
		DeckId    string `json:"deck_id"`
		Shuffled  bool   `json:"shuffled"`
		Remaining int    `json:"remaining"`
	}

	// Test without query parameters
	req, _ := http.NewRequest("GET", "/deck", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	err := json.NewDecoder(resp.Body).Decode(&resDeck)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Nil(t, err)

	var resDraw struct {
		Cards []deck.Card `json:"cards"`
	}

	// Test with valid deck ID
	req, _ = http.NewRequest("GET", "/deck/draw/"+resDeck.DeckId, nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	err = json.NewDecoder(resp.Body).Decode(&resDraw)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(resDraw.Cards))

	// Test with correct query parameters
	req, _ = http.NewRequest("GET", "/deck/draw/"+resDeck.DeckId+"?count=10", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	err = json.NewDecoder(resp.Body).Decode(&resDraw)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(resDraw.Cards))

	// Test with incorrect query parameters
	req, _ = http.NewRequest("GET", "/deck/draw/"+resDeck.DeckId+"?count=50", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	err = json.NewDecoder(resp.Body).Decode(&resDraw)
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Test with invalid deck ID
	req, _ = http.NewRequest("GET", "/deck/draw/123", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	err = json.NewDecoder(resp.Body).Decode(&resDraw)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
