package dummy

import (
	"go.fluxy.net/undeck"
	"go.fluxy.net/undeck/decks"
	"go.fluxy.net/undeck/repo"
)

// New is a factory function for a deck
func New(id string) undeck.Deck {
	return deck{
		id: id,
	}
}

// NewWithCards creates a new deck with a given id and cards
func NewWithCards(id string, isShuffled bool, cards ...undeck.Card) undeck.Deck {
	return deck{
		id:         id,
		isShuffled: isShuffled,
		cards:      cards,
	}
}

// NewDecker returns a deck factory that will always yield the same cards. useful for testing
func NewDecker(cards ...undeck.Card) repo.DeckerFunc {
	return func(id string) undeck.Deck {
		return deck{
			id:    id,
			cards: cards,
		}
	}
}

// deck is an implementation of a deck suitable for testing
type deck struct {
	id         string
	isShuffled bool
	cards      []undeck.Card
}

func (d deck) ID() string {
	return d.id
}

func (d deck) IsShuffled() bool {
	return d.isShuffled
}

func (d deck) Remaining() int {
	return len(d.cards)
}

func (d deck) Cards() []undeck.Card {
	var c []undeck.Card

	for i := range d.cards {
		c = append(c, d.cards[i].Duplicate())
	}

	return c
}

// Add cards to a deck
func (d deck) Add(cards ...undeck.Card) undeck.Deck {
	return deck{
		id:         d.id,
		isShuffled: d.isShuffled,
		cards:      decks.AddCards(d.cards, cards...),
	}
}

// Draw cards from the deck
func (d deck) Draw(count int) (undeck.Deck, []undeck.Card, error) {
	var remaining, removed, err = decks.RemoveCards(d.cards, count)

	if err != nil {
		return nil, nil, err
	}

	d.cards = remaining

	return d, removed, nil
}

// Shuffle the deck by moving first to second and second to first
func (d deck) Shuffle() undeck.Deck {
	var dup = deck{
		id:         d.id,
		isShuffled: true,
		cards:      d.Cards(),
	}

	if len(dup.cards) > 2 {
		dup.cards[0], dup.cards[1] = dup.cards[1], dup.cards[0]
	}

	return dup
}

// Duplicate the deck
func (d deck) Duplicate() undeck.Deck {
	return deck{
		id:         d.id,
		isShuffled: d.isShuffled,
		cards:      d.Cards(),
	}
}
