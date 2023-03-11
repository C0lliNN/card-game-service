// Package game implements the domain entities for managing a card game. It contains logic about Deck, Card, Suit and Rank.
//
// The game package should contain business rules, and all additional capabilities should be injected through interfaces
// like DeckRepository and IDGenerator

package game

import "context"

// DeckRepository performs deck persistence
type DeckRepository interface {
	Save(ctx context.Context, deck Deck) error
	FindByID(ctx context.Context, id string) (Deck, error)
}

// IDGenerator generates a new ID. It's very useful for testing since it's possible to mock it
type IDGenerator interface {
	NewID() string
}

// Game acts as a Facade providing to external packages an entry-point to perform high-level operations (use cases)
type Game struct {
	repo        DeckRepository
	idGenerator IDGenerator
}

// CreateDeckRequest aggregates the data necessary for creating a Deck
type CreateDeckRequest struct {
	Shuffled bool     `json:"shuffled"`
	Partial  bool     `json:"partial"`
	Cards    []string `json:"cards"`
}

// DrawCardsRequest aggregates the data necessary for drawing Card from a Deck
type DrawCardsRequest struct {
	DeckID   string
	Quantity int `json:"quantity"`
}

// DrawCardsResponse represents a response for a Draw operation
type DrawCardsResponse struct {
	Cards []CardResponse `json:"cards"`
}

// DeckResponse represents a Deck externally. Using this pattern instead of exposing the Deck directly allow modifications
// to the Deck entity without breaking external clients.
type DeckResponse struct {
	DeckID    string         `json:"deck_id"`
	Shuffled  bool           `json:"shuffled"`
	Remaining int            `json:"remaining"`
	Cards     []CardResponse `json:"cards"`
}

// CardResponse represents a Card externally. Using this pattern instead of exposing the Card directly allow modifications
// to the Card entity without breaking external clients.
type CardResponse struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

// NewGame returns a new Game
func NewGame(repo DeckRepository, idGenerator IDGenerator) Game {
	return Game{repo: repo, idGenerator: idGenerator}
}

// CreateDeck creates a new Deck based on the request. An error is returned if that cannot be done
func (g Game) CreateDeck(ctx context.Context, request CreateDeckRequest) (DeckResponse, error) {
	deck := CreateDefaultDeck(g.idGenerator.NewID())
	if request.Partial {
		deck.SwitchToPartial(request.Cards)
	}

	if request.Shuffled {
		deck.Shuffle()
	}

	if err := g.repo.Save(ctx, deck); err != nil {
		return DeckResponse{}, err
	}

	return createDeckResponseFromDeck(deck), nil
}

// OpenDeck returns an existing Deck by the id. An error is returned if that cannot be done
func (g Game) OpenDeck(ctx context.Context, deckId string) (DeckResponse, error) {
	deck, err := g.repo.FindByID(ctx, deckId)
	if err != nil {
		return DeckResponse{}, err
	}

	return createDeckResponseFromDeck(deck), nil
}

// DrawCards Removes the specified quantity of Cards from the deck and returns them. An error is returned if that cannot be done
func (g Game) DrawCards(ctx context.Context, request DrawCardsRequest) (DrawCardsResponse, error) {
	deck, err := g.repo.FindByID(ctx, request.DeckID)
	if err != nil {
		return DrawCardsResponse{}, err
	}

	cards, err := deck.Draw(request.Quantity)
	if err != nil {
		return DrawCardsResponse{}, err
	}

	if err = g.repo.Save(ctx, deck); err != nil {
		return DrawCardsResponse{}, err
	}

	return createDrawCardsResponseFromCards(cards), nil
}

func createDeckResponseFromDeck(deck Deck) DeckResponse {
	cardResponses := make([]CardResponse, len(deck.Cards))
	for i := len(deck.Cards) - 1; i >= 0; i-- {
		cardResponses[i] = createCardResponseFromCard(deck.Cards[i])
	}

	return DeckResponse{
		DeckID:    deck.ID,
		Shuffled:  deck.Shuffled,
		Remaining: deck.TotalCardsRemaining(),
		Cards:     cardResponses,
	}
}

func createDrawCardsResponseFromCards(cards []Card) DrawCardsResponse {
	cardResponses := make([]CardResponse, len(cards))
	for i := range cards {
		cardResponses[i] = createCardResponseFromCard(cards[i])
	}

	return DrawCardsResponse{Cards: cardResponses}
}

func createCardResponseFromCard(card Card) CardResponse {
	return CardResponse{
		Value: card.Rank.String(),
		Suit:  card.Suit.String(),
		Code:  card.Code(),
	}
}
