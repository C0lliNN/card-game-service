package game

import (
	"math/rand"
	"time"
)

// Deck represents a deck of cards
type Deck struct {
	ID       string `bson:"id"`
	Shuffled bool
	Partial  bool
	Cards    []Card
}

// TotalCardsRemaining returns the total count cards in a deck
func (d *Deck) TotalCardsRemaining() int {
	return len(d.Cards)
}

// SwitchToPartial turns the deck into a partial deck only containing the allowed cards
func (d *Deck) SwitchToPartial(allowedCardCodes []string) {
	newCards := make([]Card, 0)
	for _, c := range d.Cards {
		if contains(allowedCardCodes, c.Code()) {
			newCards = append(newCards, c)
		}
	}

	d.Cards = newCards
	d.Partial = true
}

// Shuffle sorts cards in a random order
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
	d.Shuffled = true
}

// Draw removes the requested card quantity from the deck and returns the cards. If the amount specified is greater than
// the available, an error is returned.
func (d *Deck) Draw(quantity int) ([]Card, error) {
	cardsRemaining := d.TotalCardsRemaining()
	if quantity < 0 || quantity > cardsRemaining {
		return nil, ErrInvalidDrawQuantity
	}

	cards := d.Cards[:quantity]
	d.Cards = d.Cards[quantity:]

	return cards, nil
}

func contains(slice []string, el string) bool {
	for i := range slice {
		if slice[i] == el {
			return true
		}
	}

	return false
}
