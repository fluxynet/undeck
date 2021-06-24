package card

import (
	"go.fluxy.net/undeck"
	"strings"
)

// FromStringer converts a string to a card
type FromStringer func(s string) (undeck.Card, error)

// FromString returns multiple cards from a comma separated list of codes
func FromString(f FromStringer, raw string) ([]undeck.Card, error) {
	var (
		err error

		raws  = strings.Split(raw, ",")
		cards = make([]undeck.Card, len(raw))
	)

	for i := range raws {
		cards[i], err = f(raws[i])
		if err != nil {
			return nil, err
		}
	}

	return cards, nil
}
