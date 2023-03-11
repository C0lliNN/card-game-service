package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeck_TotalCardsRemaining(t *testing.T) {
	d := Deck{}
	assert.Equal(t, 0, d.TotalCardsRemaining())

	d.Cards = []Card{{Rank: Ace, Suit: Spades}, {Rank: King, Suit: Hearts}}
	assert.Equal(t, 2, d.TotalCardsRemaining())
}

func TestDeck_SwitchToPartial(t *testing.T) {
	d := Deck{Cards: []Card{
		{Rank: Ace, Suit: Spades},
		{Rank: King, Suit: Hearts},
		{Rank: Rank3, Suit: Diamonds},
		{Rank: Rank6, Suit: Clubs},
		{Rank: Jack, Suit: Spades},
		{Rank: Queen, Suit: Hearts},
		{Rank: Rank7, Suit: Diamonds},
		{Rank: Rank9, Suit: Clubs},
	}}

	d.SwitchToPartial([]string{"AS", "3D", "JS"})

	assert.True(t, d.Partial)
	assert.Equal(t, []Card{{Rank: Ace, Suit: Spades}, {Rank: Rank3, Suit: Diamonds}, {Rank: Jack, Suit: Spades}}, d.Cards)
}

func TestDeck_Shuffle(t *testing.T) {
	d := Deck{Cards: []Card{
		{Rank: Ace, Suit: Spades},
		{Rank: King, Suit: Hearts},
		{Rank: Rank3, Suit: Diamonds},
		{Rank: Rank6, Suit: Clubs},
		{Rank: Jack, Suit: Spades},
		{Rank: Queen, Suit: Hearts},
		{Rank: Rank7, Suit: Diamonds},
		{Rank: Rank9, Suit: Clubs},
	}}

	d.Shuffle()

	assert.True(t, d.Shuffled)
}

func TestDeck_Draw(t *testing.T) {
	cards := []Card{
		{Rank: Ace, Suit: Spades},
		{Rank: King, Suit: Hearts},
		{Rank: Rank3, Suit: Diamonds},
	}

	tests := []struct {
		Name                   string
		Quantity               int
		ExpectedCards          []Card
		ExpectedCardsRemaining int
		ExpectedErr            error
	}{
		{
			Name:                   "Negative quantity",
			Quantity:               -1,
			ExpectedCards:          nil,
			ExpectedCardsRemaining: 3,
			ExpectedErr:            ErrInvalidDrawQuantity,
		},
		{
			Name:                   "Quantity greater than available size",
			Quantity:               4,
			ExpectedCards:          nil,
			ExpectedCardsRemaining: 3,
			ExpectedErr:            ErrInvalidDrawQuantity,
		},
		{
			Name:                   "Quantity equal to the available size",
			Quantity:               3,
			ExpectedCards:          cards,
			ExpectedCardsRemaining: 0,
			ExpectedErr:            nil,
		},
		{
			Name:                   "Quantity smaller than the available size",
			Quantity:               1,
			ExpectedCards:          cards[2:3],
			ExpectedCardsRemaining: 2,
			ExpectedErr:            nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			d := Deck{Cards: cards}

			actualCards, err := d.Draw(tc.Quantity)

			assert.Equal(t, tc.ExpectedErr, err)
			assert.Equal(t, tc.ExpectedCards, actualCards)
			assert.Equal(t, tc.ExpectedCardsRemaining, d.TotalCardsRemaining())
		})
	}
}
