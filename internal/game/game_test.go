package game

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GameTestSuite struct {
	suite.Suite
	repo        *DeckRepositoryMock
	idGenerator *IDGeneratorMock
	game        Game
}

func (s *GameTestSuite) SetupTest() {
	s.repo = new(DeckRepositoryMock)
	s.idGenerator = new(IDGeneratorMock)
	s.game = NewGame(s.repo, s.idGenerator)
}

func TestGame(t *testing.T) {
	suite.Run(t, new(GameTestSuite))
}

func (s *GameTestSuite) TestCreateDeck_SaveFails() {
	request := CreateDeckRequest{Shuffled: false, Partial: false, Cards: nil}
	s.idGenerator.On("NewID").Return("deck-id")
	s.repo.On("Save", mock.Anything, mock.Anything).Return(fmt.Errorf("some error"))

	_, err := s.game.CreateDeck(context.TODO(), request)

	assert.Equal(s.T(), fmt.Errorf("some error"), err)
	s.repo.AssertNumberOfCalls(s.T(), "Save", 1)
	s.idGenerator.AssertNumberOfCalls(s.T(), "NewID", 1)
}

func (s *GameTestSuite) TestCreateDeck_FullNotShuffledDeck() {
	request := CreateDeckRequest{Shuffled: false, Partial: false, Cards: nil}
	s.idGenerator.On("NewID").Return("deck-id")

	expectedDeck := CreateDefaultDeck("deck-id")
	s.repo.On("Save", context.TODO(), expectedDeck).Return(nil)

	expectedDeckResponse := createDeckResponseFromDeck(expectedDeck)
	deckResponse, err := s.game.CreateDeck(context.TODO(), request)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedDeckResponse, deckResponse)

	s.repo.AssertNumberOfCalls(s.T(), "Save", 1)
	s.idGenerator.AssertNumberOfCalls(s.T(), "NewID", 1)
}

func (s *GameTestSuite) TestCreateDeck_FullShuffledDeck() {
	request := CreateDeckRequest{Shuffled: true, Partial: false, Cards: nil}
	s.idGenerator.On("NewID").Return("deck-id")
	s.repo.On("Save", context.TODO(), mock.Anything).Return(nil)

	deckResponse, err := s.game.CreateDeck(context.TODO(), request)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "deck-id", deckResponse.DeckID)
	assert.True(s.T(), deckResponse.Shuffled)
	assert.Len(s.T(), deckResponse.Cards, 52)

	s.repo.AssertNumberOfCalls(s.T(), "Save", 1)
	s.idGenerator.AssertNumberOfCalls(s.T(), "NewID", 1)
}

func (s *GameTestSuite) TestCreateDeck_PartialNotShuffledDeck() {
	request := CreateDeckRequest{Shuffled: false, Partial: true, Cards: []string{"AS", "4D", "KH"}}
	s.idGenerator.On("NewID").Return("deck-id")
	s.repo.On("Save", context.TODO(), mock.Anything).Return(nil)

	deckResponse, err := s.game.CreateDeck(context.TODO(), request)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "deck-id", deckResponse.DeckID)
	assert.False(s.T(), deckResponse.Shuffled)
	assert.Len(s.T(), deckResponse.Cards, 3)

	s.repo.AssertNumberOfCalls(s.T(), "Save", 1)
	s.idGenerator.AssertNumberOfCalls(s.T(), "NewID", 1)
}

func (s *GameTestSuite) TestOpenDeck_DeckNotFound() {
	id := "deck-id"
	s.repo.On("FindByID", context.TODO(), id).Return(Deck{}, ErrDeckNotFound)

	_, err := s.game.OpenDeck(context.TODO(), id)

	assert.Equal(s.T(), ErrDeckNotFound, err)
	s.repo.AssertNumberOfCalls(s.T(), "FindByID", 1)
}

func (s *GameTestSuite) TestOpenDeck_DeckFound() {
	deck := CreateDefaultDeck("deck-id")
	s.repo.On("FindByID", context.TODO(), deck.ID).Return(deck, nil)

	deckResponse, err := s.game.OpenDeck(context.TODO(), deck.ID)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), createDeckResponseFromDeck(deck), deckResponse)
	s.repo.AssertNumberOfCalls(s.T(), "FindByID", 1)
}

func (s *GameTestSuite) TestDrawCards_DeckNotFound() {
	req := DrawCardsRequest{DeckID: "deck-id", Quantity: 5}
	s.repo.On("FindByID", context.TODO(), req.DeckID).Return(Deck{}, ErrDeckNotFound)

	_, err := s.game.DrawCards(context.TODO(), req)

	assert.Equal(s.T(), ErrDeckNotFound, err)
	s.repo.AssertNumberOfCalls(s.T(), "FindByID", 1)
	s.repo.AssertNumberOfCalls(s.T(), "Save", 0)
}

func (s *GameTestSuite) TestDrawCards_InvalidQuantity() {
	req := DrawCardsRequest{DeckID: "deck-id", Quantity: 55}
	deck := CreateDefaultDeck(req.DeckID)
	s.repo.On("FindByID", context.TODO(), req.DeckID).Return(deck, nil)

	_, err := s.game.DrawCards(context.TODO(), req)

	assert.Equal(s.T(), ErrInvalidDrawQuantity, err)
	s.repo.AssertNumberOfCalls(s.T(), "FindByID", 1)
	s.repo.AssertNumberOfCalls(s.T(), "Save", 0)
}

func (s *GameTestSuite) TestDrawCards_FailOnSave() {
	req := DrawCardsRequest{DeckID: "deck-id", Quantity: 5}
	deck := CreateDefaultDeck(req.DeckID)
	s.repo.On("FindByID", context.TODO(), req.DeckID).Return(deck, nil)
	s.repo.On("Save", context.TODO(), mock.Anything).Return(fmt.Errorf("some error"))

	_, err := s.game.DrawCards(context.TODO(), req)

	assert.Equal(s.T(), fmt.Errorf("some error"), err)
	s.repo.AssertNumberOfCalls(s.T(), "FindByID", 1)
	s.repo.AssertNumberOfCalls(s.T(), "Save", 1)
}

func (s *GameTestSuite) TestDrawCards_Success() {
	req := DrawCardsRequest{DeckID: "deck-id", Quantity: 5}
	deck := CreateDefaultDeck(req.DeckID)
	s.repo.On("FindByID", context.TODO(), req.DeckID).Return(deck, nil)
	s.repo.On("Save", context.TODO(), mock.Anything).Return(nil)

	drawResponse, err := s.game.DrawCards(context.TODO(), req)

	assert.NoError(s.T(), err)
	assert.Len(s.T(), drawResponse.Cards, 5)
	s.repo.AssertNumberOfCalls(s.T(), "FindByID", 1)
	s.repo.AssertNumberOfCalls(s.T(), "Save", 1)
}
