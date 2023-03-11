package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateDefaultDeck(t *testing.T) {
	id := "deck-id"

	expectedDeck := Deck{
		ID:       id,
		Shuffled: false,
		Partial:  false,
		Cards: []Card{
			{Suit: Spades, Rank: Ace},
			{Suit: Spades, Rank: Rank2},
			{Suit: Spades, Rank: Rank3},
			{Suit: Spades, Rank: Rank4},
			{Suit: Spades, Rank: Rank5},
			{Suit: Spades, Rank: Rank6},
			{Suit: Spades, Rank: Rank7},
			{Suit: Spades, Rank: Rank8},
			{Suit: Spades, Rank: Rank9},
			{Suit: Spades, Rank: Rank10},
			{Suit: Spades, Rank: Jack},
			{Suit: Spades, Rank: Queen},
			{Suit: Spades, Rank: King},
			{Suit: Diamonds, Rank: Ace},
			{Suit: Diamonds, Rank: Rank2},
			{Suit: Diamonds, Rank: Rank3},
			{Suit: Diamonds, Rank: Rank4},
			{Suit: Diamonds, Rank: Rank5},
			{Suit: Diamonds, Rank: Rank6},
			{Suit: Diamonds, Rank: Rank7},
			{Suit: Diamonds, Rank: Rank8},
			{Suit: Diamonds, Rank: Rank9},
			{Suit: Diamonds, Rank: Rank10},
			{Suit: Diamonds, Rank: Jack},
			{Suit: Diamonds, Rank: Queen},
			{Suit: Diamonds, Rank: King},
			{Suit: Clubs, Rank: Ace},
			{Suit: Clubs, Rank: Rank2},
			{Suit: Clubs, Rank: Rank3},
			{Suit: Clubs, Rank: Rank4},
			{Suit: Clubs, Rank: Rank5},
			{Suit: Clubs, Rank: Rank6},
			{Suit: Clubs, Rank: Rank7},
			{Suit: Clubs, Rank: Rank8},
			{Suit: Clubs, Rank: Rank9},
			{Suit: Clubs, Rank: Rank10},
			{Suit: Clubs, Rank: Jack},
			{Suit: Clubs, Rank: Queen},
			{Suit: Clubs, Rank: King},
			{Suit: Hearts, Rank: Ace},
			{Suit: Hearts, Rank: Rank2},
			{Suit: Hearts, Rank: Rank3},
			{Suit: Hearts, Rank: Rank4},
			{Suit: Hearts, Rank: Rank5},
			{Suit: Hearts, Rank: Rank6},
			{Suit: Hearts, Rank: Rank7},
			{Suit: Hearts, Rank: Rank8},
			{Suit: Hearts, Rank: Rank9},
			{Suit: Hearts, Rank: Rank10},
			{Suit: Hearts, Rank: Jack},
			{Suit: Hearts, Rank: Queen},
			{Suit: Hearts, Rank: King},
		},
	}

	actualDeck := CreateDefaultDeck(id)

	assert.Equal(t, expectedDeck, actualDeck)
}
