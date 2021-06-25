package decks

import "go.fluxy.net/undeck"

// AddCards to a slice of cards
func AddCards(cards []undeck.Card, more ...undeck.Card) []undeck.Card {
	var c []undeck.Card

	for i := range cards {
		c = append(c, cards[i].Duplicate())
	}

	for i := range more {
		c = append(c, more[i].Duplicate())
	}

	return c
}

// RemoveCards from a slice returning remaining cards and removed cards
func RemoveCards(cards []undeck.Card, count int) (remaining, removed []undeck.Card, err error) {
	if len(cards) < count {
		return cards, nil, undeck.ErrNotEnoughCards
	}

	var c []undeck.Card

	for i := 0; i < count; i++ {
		c = append(c, cards[i].Duplicate())
	}

	cards = cards[count:]

	return cards, c, nil
}
