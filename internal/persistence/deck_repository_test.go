package persistence_test

import (
	"C0lliNN/card-game-service/internal/game"
	"C0lliNN/card-game-service/internal/persistence"
	"C0lliNN/card-game-service/internal/test"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DeckRepositoryTestSuite struct {
	test.RepositoryTestSuite
	repo persistence.DeckRepository
}

func (s *DeckRepositoryTestSuite) SetupSuite() {
	s.RepositoryTestSuite.SetupSuite()

	s.repo = persistence.NewDeckRepository(s.DB)
}

func TestDeckRepository(t *testing.T) {
	suite.Run(t, new(DeckRepositoryTestSuite))
}

func (s *DeckRepositoryTestSuite) TestSave_Insert() {
	d := game.Deck{
		ID:       "my-id",
		Shuffled: false,
		Partial:  false,
		Cards: []game.Card{
			{Rank: game.Ace, Suit: game.Spades},
			{Rank: game.King, Suit: game.Hearts},
			{Rank: game.Rank3, Suit: game.Diamonds},
			{Rank: game.Rank6, Suit: game.Clubs},
			{Rank: game.Jack, Suit: game.Spades},
			{Rank: game.Queen, Suit: game.Hearts},
			{Rank: game.Rank7, Suit: game.Diamonds},
			{Rank: game.Rank9, Suit: game.Clubs},
		},
	}

	ctx := context.Background()

	err := s.repo.Save(ctx, d)
	require.NoError(s.T(), err)

	found, err := s.repo.FindByID(ctx, d.ID)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), d, found)
}

func (s *DeckRepositoryTestSuite) TestSave_Update() {
	d := game.Deck{
		ID:       "my-id",
		Shuffled: false,
		Partial:  false,
		Cards: []game.Card{
			{Rank: game.Ace, Suit: game.Spades},
			{Rank: game.King, Suit: game.Hearts},
			{Rank: game.Rank3, Suit: game.Diamonds},
			{Rank: game.Rank6, Suit: game.Clubs},
			{Rank: game.Jack, Suit: game.Spades},
			{Rank: game.Queen, Suit: game.Hearts},
			{Rank: game.Rank7, Suit: game.Diamonds},
			{Rank: game.Rank9, Suit: game.Clubs},
		},
	}

	ctx := context.Background()

	err := s.repo.Save(ctx, d)
	require.NoError(s.T(), err)

	d.Shuffle()
	err = s.repo.Save(ctx, d)
	require.NoError(s.T(), err)

	found, err := s.repo.FindByID(ctx, d.ID)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), d, found)
}

func (s *DeckRepositoryTestSuite) TestFindByID_NotFound() {
	ctx := context.TODO()
	_, err := s.repo.FindByID(ctx, "some-id")
	assert.Equal(s.T(), game.ErrDeckNotFound, err)
}

func (s *DeckRepositoryTestSuite) TestFindByID_Found() {
	d := game.Deck{
		ID:       "my-id",
		Shuffled: false,
		Partial:  false,
		Cards: []game.Card{
			{Rank: game.Ace, Suit: game.Spades},
			{Rank: game.King, Suit: game.Hearts},
			{Rank: game.Rank3, Suit: game.Diamonds},
			{Rank: game.Rank6, Suit: game.Clubs},
			{Rank: game.Jack, Suit: game.Spades},
			{Rank: game.Queen, Suit: game.Hearts},
			{Rank: game.Rank7, Suit: game.Diamonds},
			{Rank: game.Rank9, Suit: game.Clubs},
		},
	}

	ctx := context.Background()

	err := s.repo.Save(ctx, d)
	require.NoError(s.T(), err)

	found, err := s.repo.FindByID(ctx, d.ID)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), d, found)
}
