// go:generate stringer -type=Suit -linecomment

package game

// Suit represents a card suit. The stringer go tool is being used for generating the String method that returns the
// uppercase suit name
type Suit int

// Code returns n unique identifier for the Suit
func (s Suit) Code() string {
	suit := s.String()
	return string(suit[0])
}

const (
	Spades   Suit = iota + 1 // SPADES
	Diamonds                 // DIAMONDS
	Clubs                    // CLUBS
	Hearts                   // HEARTS
)
