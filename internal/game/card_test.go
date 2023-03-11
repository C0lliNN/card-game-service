package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCard_Code(t *testing.T) {
	tests := []struct {
		Name     string
		Rank     Rank
		Suit     Suit
		Expected string
	}{
		{
			Name:     "Ace of Spades",
			Rank:     Ace,
			Suit:     Spades,
			Expected: "AS",
		},
		{
			Name:     "Three of Diamonds",
			Rank:     Rank3,
			Suit:     Diamonds,
			Expected: "3D",
		},
		{
			Name:     "Jack of Clubs",
			Rank:     Jack,
			Suit:     Clubs,
			Expected: "JC",
		},
		{
			Name:     "King of Hearts",
			Rank:     King,
			Suit:     Hearts,
			Expected: "KH",
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			c := Card{Suit: tc.Suit, Rank: tc.Rank}
			assert.Equal(t, tc.Expected, c.Code())
		})
	}
}
