package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuit_Code(t *testing.T) {
	tests := []struct {
		Name     string
		Suit     Suit
		Expected string
	}{
		{
			Name:     "Spades",
			Suit:     Spades,
			Expected: "S",
		},
		{
			Name:     "Diamonds",
			Suit:     Diamonds,
			Expected: "D",
		},
		{
			Name:     "Clubs",
			Suit:     Clubs,
			Expected: "C",
		},
		{
			Name:     "Hearts",
			Suit:     Hearts,
			Expected: "H",
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Expected, tc.Suit.Code())
		})
	}
}
