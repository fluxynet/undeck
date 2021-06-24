package memory

import (
	"context"
	"github.com/google/uuid"
	"go.fluxy.net/undeck"
)

// New repo instance
func New() undeck.Repo {
	return &Repo{
		decks: make(map[string]undeck.Deck),
	}
}

// Repo for in memory persistence
type Repo struct {
	decks map[string]undeck.Deck
}

func (r *Repo) Save(ctx context.Context, deck undeck.Deck) (undeck.Deck, error) {
	if deck.ID != "" {
		// already have an id yay
	} else if id, err := uuid.NewRandom(); err != nil {
		return deck, err
	} else {
		deck.ID = id.String()
	}

	r.decks[deck.ID] = deck

	return deck, nil
}

func (r *Repo) Find(ctx context.Context, id string) (*undeck.Deck, error) {
	var (
		d, ok = r.decks[id]
		d2    undeck.Deck
	)

	if !ok {
		return nil, undeck.ErrDeckNotFound
	}

	d2.ID = d.ID
	d2.IsShuffled = d.IsShuffled
	d2.Cards = make([]undeck.Card, len(d.Cards))

	for i := range d.Cards {
		d2.Cards[i] = d.Cards[i].Duplicate()
	}

	return &d2, nil
}
