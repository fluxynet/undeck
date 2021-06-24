package undeck

import (
	"math/rand"
	"time"
)

type Deck struct {
	ID         string
	IsShuffled bool
	Cards      []Card
}

func (d Deck) Remaining() int {
	return len(d.Cards)
}

func (d Deck) Add(cards ...Card) Deck {
	var (
		t = len(d.Cards)
		c = make([]Card, len(cards)+t)
	)

	for i := range d.Cards {
		c[i] = d.Cards[i].Duplicate()
	}

	for i := range cards {
		c[t+i] = cards[i].Duplicate()
	}

	return Deck{
		ID:         d.ID,
		IsShuffled: d.IsShuffled,
		Cards:      c,
	}
}

func (d Deck) Draw(count int) (Deck, []Card, error) {
	var t = len(d.Cards)

	if t < count {
		return d, nil, ErrNotEnoughCards
	}

	var c = make([]Card, count)

	for i := range c {
		c[i] = d.Cards[i].Duplicate()
	}

	d.Cards = d.Cards[count:]

	return d, c, nil
}

func (d Deck) Shuffle() Deck {
	var t = len(d.Cards)

	if t == 0 {
		return d
	}

	var (
		c = make([]Card, t)
	)

	for i := range d.Cards {
		c[i] = d.Cards[i].Duplicate()
	}

	d.Cards = c
	d.IsShuffled = true

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(t, func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})

	return d
}
