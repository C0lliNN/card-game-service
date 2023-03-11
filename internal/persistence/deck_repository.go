// Package persistence it implements persistence operations. The code is in this package should not contain any
// business rules

package persistence

import (
	"C0lliNN/card-game-service/internal/game"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	deckCollection = "deck"
)

// DeckRepository performs persistence operations with the game.Deck
type DeckRepository struct {
	db *mongo.Database
}

func NewDeckRepository(db *mongo.Database) DeckRepository {
	return DeckRepository{db: db}
}

// Save inserts a new deck if not present or update an existing deck
func (r DeckRepository) Save(ctx context.Context, deck game.Deck) error {
	log.Println("Saving deck with id: ", deck.ID)
	_, err := r.db.Collection(deckCollection).ReplaceOne(ctx, bson.M{"_id": deck.ID}, deck, options.Replace().SetUpsert(true))
	return err
}

// FindByID tries to find a deck for the given id
func (r DeckRepository) FindByID(ctx context.Context, id string) (game.Deck, error) {
	log.Println("Retrieving new deck with id: ", id)

	var deck game.Deck

	if err := r.db.Collection(deckCollection).FindOne(ctx, bson.M{"_id": id}).Decode(&deck); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return game.Deck{}, game.ErrDeckNotFound
		}
		return game.Deck{}, err
	}

	return deck, nil
}
