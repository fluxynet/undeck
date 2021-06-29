package memory

import (
	"context"
	"github.com/google/uuid"
	"go.fluxy.net/undeck"
	"go.fluxy.net/undeck/repo"
)

// New repo instance
func New() undeck.Repo {
	return &Repo{
		decks: make(map[string]undeck.Deck),
	}
}

// NewWith instantiates a repo with custom decks and id generator
func NewWith(fi repo.IDGenerator, shufflerFunc undeck.ShufflerFunc, decks ...undeck.Deck) undeck.Repo {
	var d = make(map[string]undeck.Deck, len(decks))

	for i := range decks {
		d[decks[i].ID] = decks[i]
	}

	return &Repo{
		decks:        d,
		idGenerator:  fi,
		shufflerFunc: shufflerFunc,
	}
}

// Repo for in memory persistence
type Repo struct {
	decks        map[string]undeck.Deck
	idGenerator  repo.IDGenerator
	shufflerFunc undeck.ShufflerFunc
}

func (r *Repo) Create(ctx context.Context) (undeck.Deck, error) {
	var d undeck.Deck

	if r.idGenerator != nil {
		d.ID = r.idGenerator()
	} else if u, err := uuid.NewRandom(); err != nil {
		return d, err
	} else {
		d.ID = u.String()
	}

	d.Shuffler = r.shufflerFunc

	return d, nil
}

func (r *Repo) Save(ctx context.Context, deck undeck.Deck) (undeck.Deck, error) {
	r.decks[deck.ID] = deck

	return deck, nil
}

func (r *Repo) Find(ctx context.Context, id string) (undeck.Deck, error) {
	var (
		d, ok = r.decks[id]
		d2    undeck.Deck
	)

	if !ok {
		return d2, undeck.ErrDeckNotFound
	}

	d2 = d.Duplicate()

	return d2, nil
}

func (r Repo) Dump() map[string]undeck.Deck {
	return r.decks
}
