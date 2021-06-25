package decks

import (
	"go.fluxy.net/undeck"
	"go.fluxy.net/undeck/decks"
	"math/rand"
	"time"
)

// New deck
func New(id string) undeck.Deck {
	return deck{
		id:         id,
		isShuffled: false,
		cards:      nil,
	}
}

// deck is an implementation of a deck suitable for most cases
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

func (d deck) Add(cards ...undeck.Card) undeck.Deck {
	return deck{
		id:         d.id,
		isShuffled: d.isShuffled,
		cards:      decks.AddCards(d.cards, cards...),
	}
}

func (d deck) Draw(count int) (undeck.Deck, []undeck.Card, error) {
	var remaining, removed, err = decks.RemoveCards(d.cards, count)

	if err != nil {
		return nil, nil, err
	}

	d.cards = remaining

	return d, removed, nil
}

func (d deck) Shuffle() undeck.Deck {
	if len(d.cards) == 0 {
		return d
	}

	d.cards = d.Cards()
	d.isShuffled = true

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(d.Remaining(), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})

	return d
}

func (d deck) Duplicate() undeck.Deck {
	return deck{
		id:         d.id,
		isShuffled: d.isShuffled,
		cards:      d.Cards(),
	}
}

func (d deck) Cards() []undeck.Card {
	var cards []undeck.Card

	for i := range d.cards {
		cards = append(cards, d.cards[i].Duplicate())
	}

	return cards
}
