package server

import (
	"C0lliNN/card-game-service/internal/game"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GameHandler declares the routes and handlers allowing the clients to interact with the game module through HTTP
type GameHandler struct {
	Game game.Game
}

func NewGameHandler(g game.Game) GameHandler {
	return GameHandler{Game: g}
}

func (h GameHandler) Routes() []Route {
	return []Route{
		{Method: http.MethodPost, Path: "/decks", Handler: h.createDeck},
		{Method: http.MethodGet, Path: "/decks/:id", Handler: h.openDeck},
		// Since this operation is not idempotent (cards are removed from the deck on each request),
		// the PATCH verb seems like a good fit since this endpoint will apply a partial update in the deck resource
		{Method: http.MethodPatch, Path: "/decks/:id/draw", Handler: h.drawCards},
	}
}

func (h GameHandler) createDeck(c *gin.Context) {
	log.Println("New request for creating a deck received")

	var req game.CreateDeckRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	response, err := h.Game.CreateDeck(c, req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h GameHandler) openDeck(c *gin.Context) {
	log.Println("New request for opening a deck received")

	response, err := h.Game.OpenDeck(c, c.Param("id"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h GameHandler) drawCards(c *gin.Context) {
	log.Println("New request for drawing cards from a deck recieved")

	var req game.DrawCardsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	req.DeckID = c.Param("id")

	response, err := h.Game.DrawCards(c, req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}
