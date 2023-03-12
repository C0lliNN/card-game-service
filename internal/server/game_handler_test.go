package server_test

import (
	"C0lliNN/card-game-service/internal/game"
	"C0lliNN/card-game-service/internal/test"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/h2non/baloo.v3"
	"net/http"
	"testing"
)

type GameHandlerTestSuite struct {
	test.ServerTestSuite
	balooClient *baloo.Client
}

func TestGameHandler(t *testing.T) {
	suite.Run(t, new(GameHandlerTestSuite))
}

func (s *GameHandlerTestSuite) SetupSuite() {
	s.ServerTestSuite.SetupSuite()

	s.balooClient = baloo.New(s.BaseURL)
}

func (s *GameHandlerTestSuite) TestCreateDeck() {
	req := game.CreateDeckRequest{}
	deckResponse := game.DeckResponse{}

	err := s.balooClient.Post("/decks").
		JSON(req).
		Expect(s.T()).
		Status(http.StatusCreated).
		AssertFunc(func(response *http.Response, request *http.Request) error {
			return json.NewDecoder(response.Body).Decode(&deckResponse)
		}).
		Done()

	require.NoError(s.T(), err)
	assert.False(s.T(), deckResponse.Shuffled)
	assert.Equal(s.T(), 52, deckResponse.Remaining)
	assert.Len(s.T(), deckResponse.Cards, 52)
}

func (s *GameHandlerTestSuite) TestGetDeck() {
	req := game.CreateDeckRequest{}
	expectedDeckResponse := game.DeckResponse{}
	actualDeckResponse := game.DeckResponse{}

	err := s.balooClient.Post("/decks").
		JSON(req).
		Expect(s.T()).
		Status(http.StatusCreated).
		AssertFunc(func(response *http.Response, request *http.Request) error {
			return json.NewDecoder(response.Body).Decode(&expectedDeckResponse)
		}).
		Done()

	require.NoError(s.T(), err)

	err = s.balooClient.Get("/decks/" + expectedDeckResponse.DeckID).
		Expect(s.T()).
		Status(http.StatusOK).
		AssertFunc(func(response *http.Response, request *http.Request) error {
			return json.NewDecoder(response.Body).Decode(&actualDeckResponse)
		}).
		Done()

	require.NoError(s.T(), err)
	assert.Equal(s.T(), expectedDeckResponse, actualDeckResponse)
}

func (s *GameHandlerTestSuite) TestDrawCards() {
	createRequest := game.CreateDeckRequest{}
	deckResponse := game.DeckResponse{}
	drawResponse := game.DrawCardsResponse{}

	err := s.balooClient.Post("/decks").
		JSON(createRequest).
		Expect(s.T()).
		Status(http.StatusCreated).
		AssertFunc(func(response *http.Response, request *http.Request) error {
			return json.NewDecoder(response.Body).Decode(&deckResponse)
		}).
		Done()

	require.NoError(s.T(), err)

	drawRequest := game.DrawCardsRequest{Quantity: 3}
	err = s.balooClient.Patch("/decks/" + deckResponse.DeckID + "/draw").
		JSON(drawRequest).
		Expect(s.T()).
		Status(http.StatusOK).
		AssertFunc(func(response *http.Response, request *http.Request) error {
			return json.NewDecoder(response.Body).Decode(&drawResponse)
		}).
		Done()

	require.NoError(s.T(), err)
	assert.Len(s.T(), drawResponse.Cards, 3)
}
