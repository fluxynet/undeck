package undeck

import (
	"context"
	"errors"
)

var (
	// ErrDeckNotFound in case a deck does not exist
	ErrDeckNotFound = errors.New("deck not found")
)

// Repo is for deck persistence
type Repo interface {
	Create(ctx context.Context) (Deck, error)

	// Save a deck
	Save(ctx context.Context, deck Deck) (Deck, error)

	// Find a deck by id
	Find(ctx context.Context, id string) (Deck, error)
}
