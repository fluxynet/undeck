package undeck

import (
	"math/rand"
	"time"
)

// ShufflerFunc is a function that returns a shuffled copy of a deck
type ShufflerFunc func(Deck) Deck

func RandomShuffler(d Deck) Deck {
	if len(d.cards) == 0 {
		return d
	}

	d.cards = d.Cards()
	d.IsShuffled = true

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(d.Remaining(), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})

	return d
}

// OneTwoSwapShuffler the deck by moving first to second and second to first
func OneTwoSwapShuffler(d Deck) Deck {
	var dup = Deck{
		ID:         d.ID,
		IsShuffled: true,
		cards:      d.Cards(),
	}

	if len(dup.cards) > 2 {
		dup.cards[0], dup.cards[1] = dup.cards[1], dup.cards[0]
	}

	return dup
}

// Deck is an implementation of a deck suitable for most cases
type Deck struct {
	ID         string
	IsShuffled bool
	Shuffler   ShufflerFunc
	cards      []Card
}

func (d Deck) Remaining() int {
	return len(d.cards)
}

// Add to a slice of cards
func (d Deck) Add(cards ...Card) Deck {
	var c []Card

	for i := range d.cards {
		c = append(c, d.cards[i].Duplicate())
	}

	for i := range cards {
		c = append(c, cards[i].Duplicate())
	}

	d.cards = c
	return d
}

// Draw from a slice returning deck with remaining cards and removed cards
func (d Deck) Draw(count int) (Deck, []Card, error) {
	if len(d.cards) < count {
		return d, nil, ErrNotEnoughCards
	}

	var c []Card

	for i := 0; i < count; i++ {
		c = append(c, d.cards[i].Duplicate())
	}

	d.cards = d.cards[count:]

	return d, c, nil
}

func (d Deck) Shuffle() Deck {
	if d.Shuffler == nil {
		d.Shuffler = RandomShuffler
	}

	return d.Shuffler(d)
}

func (d Deck) Duplicate() Deck {
	return Deck{
		ID:         d.ID,
		IsShuffled: d.IsShuffled,
		cards:      d.Cards(),
	}
}

func (d Deck) Cards() []Card {
	var cards []Card

	for i := range d.cards {
		cards = append(cards, d.cards[i].Duplicate())
	}

	return cards
}
