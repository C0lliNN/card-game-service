package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRank_Code(t *testing.T) {
	tests := []struct {
		Name     string
		Rank     Rank
		Expected string
	}{
		{
			Name:     "ACE",
			Rank:     Ace,
			Expected: "A",
		},
		{
			Name:     "2",
			Rank:     Rank2,
			Expected: "2",
		},
		{
			Name:     "3",
			Rank:     Rank3,
			Expected: "3",
		},
		{
			Name:     "4",
			Rank:     Rank4,
			Expected: "4",
		},
		{
			Name:     "5",
			Rank:     Rank5,
			Expected: "5",
		},
		{
			Name:     "6",
			Rank:     Rank6,
			Expected: "6",
		},
		{
			Name:     "7",
			Rank:     Rank7,
			Expected: "7",
		},
		{
			Name:     "8",
			Rank:     Rank8,
			Expected: "8",
		},
		{
			Name:     "9",
			Rank:     Rank9,
			Expected: "9",
		},
		{
			Name:     "10",
			Rank:     Rank10,
			Expected: "10",
		},
		{
			Name:     "JACK",
			Rank:     Jack,
			Expected: "J",
		},
		{
			Name:     "QUEEN",
			Rank:     Queen,
			Expected: "Q",
		},
		{
			Name:     "KING",
			Rank:     King,
			Expected: "K",
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Expected, tc.Rank.Code())
		})
	}
}
