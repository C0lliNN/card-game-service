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
		// Although this operation is not idempotent, the DELETE verb seems applicable in this scenario since
		// Cards are going to be deleted from the deck resource.
		{Method: http.MethodDelete, Path: "/decks/:id/cards", Handler: h.drawCards},
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
	log.Println("New request for drawing cards from a deck received")

	var req game.DrawCardsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
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
