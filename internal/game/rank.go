// go:generate stringer -type=Rank -linecomment

package game

import "unicode"

// Rank represents a card rank. It goes from Ace, Rank2 up to the King. The stringer go tool is being used for generating
// the String method that returns the uppercase rank name
type Rank int

// Code returns a unique identifier for the Rank
func (r Rank) Code() string {
	rank := r.String()
	if unicode.IsNumber(rune(rank[0])) {
		return rank
	}

	return string(rank[0])
}

const (
	Ace    Rank = iota + 1 // ACE
	Rank2                  // 2
	Rank3                  // 3
	Rank4                  // 4
	Rank5                  // 5
	Rank6                  // 6
	Rank7                  // 7
	Rank8                  // 8
	Rank9                  // 9
	Rank10                 // 10
	Jack                   // JACK
	Queen                  // QUEEN
	King                   // KING
)
