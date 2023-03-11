package game

// CreateDefaultDeck given an id, returns a default Deck with 52 cards sorted
func CreateDefaultDeck(id string) Deck {
	d := Deck{ID: id, Shuffled: false, Partial: false, Cards: make([]Card, 0)}

	suits := []Suit{Spades, Diamonds, Clubs, Hearts}
	ranks := []Rank{
		Ace,
		Rank2,
		Rank3,
		Rank4,
		Rank5,
		Rank6,
		Rank7,
		Rank8,
		Rank9,
		Rank10,
		Jack,
		Queen,
		King,
	}

	for _, s := range suits {
		for _, r := range ranks {
			d.Cards = append(d.Cards, Card{Suit: s, Rank: r})
		}
	}

	return d
}
