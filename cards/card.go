package cards

import (
	"go.fluxy.net/undeck"
	"log"
	"strings"
)

// FromStringer converts a string to a card
type FromStringer func(s string) (undeck.Card, error)

// FromString returns multiple cards from a comma separated list of codes
func FromString(f FromStringer, raw string) ([]undeck.Card, error) {
	var (
		err   error
		cards []undeck.Card

		raws = strings.Split(raw, ",")
	)

	for i := range raws {
		var card undeck.Card

		card, err = f(raws[i])
		if err != nil {
			return nil, err
		}

		cards = append(cards, card)
	}

	return cards, nil
}

// MustString is FromString or fatal
func MustString(f FromStringer, raw string) []undeck.Card {
	var c, err = FromString(f, raw)

	if err != nil {
		log.Fatalln(err)
	}

	return c
}
