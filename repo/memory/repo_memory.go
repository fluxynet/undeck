package memory

import (
	"context"
	"github.com/google/uuid"
	"go.fluxy.net/undeck"
	"go.fluxy.net/undeck/repo"
)

// New repo instance
func New(f repo.DeckerFunc) undeck.Repo {
	return &Repo{
		decks:      make(map[string]undeck.Deck),
		deckerFunc: f,
	}
}

// NewWith instantiates a repo with custom decks and id generator
func NewWith(f repo.DeckerFunc, fi repo.IDGenerator, decks ...undeck.Deck) undeck.Repo {
	var d = make(map[string]undeck.Deck, len(decks))

	for i := range decks {
		d[decks[i].ID()] = decks[i]
	}

	return &Repo{
		decks:       d,
		deckerFunc:  f,
		idGenerator: fi,
	}
}

// Repo for in memory persistence
type Repo struct {
	decks       map[string]undeck.Deck
	deckerFunc  repo.DeckerFunc
	idGenerator repo.IDGenerator
}

func (r *Repo) Create(ctx context.Context) (undeck.Deck, error) {
	var id string

	if r.idGenerator != nil {
		id = r.idGenerator()
	} else if u, err := uuid.NewRandom(); err != nil {
		return nil, err
	} else {
		id = u.String()
	}

	return r.deckerFunc(id), nil
}

func (r *Repo) Save(ctx context.Context, deck undeck.Deck) (undeck.Deck, error) {
	r.decks[deck.ID()] = deck

	return deck, nil
}

func (r *Repo) Find(ctx context.Context, id string) (undeck.Deck, error) {
	var (
		d, ok = r.decks[id]
		d2    undeck.Deck
	)

	if !ok {
		return nil, undeck.ErrDeckNotFound
	}

	d2 = d.Duplicate()

	return d2, nil
}

func (r Repo) Dump() map[string]undeck.Deck {
	return r.decks
}
