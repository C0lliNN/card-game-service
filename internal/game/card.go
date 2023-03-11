package game

// Card represents a card in a Deck
type Card struct {
	Rank Rank
	Suit Suit
}

// Code returns a unique identifier for the card
func (c Card) Code() string {
	return c.Rank.Code() + c.Suit.Code()
}
