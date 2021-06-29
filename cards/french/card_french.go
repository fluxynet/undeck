package french

import (
	"go.fluxy.net/undeck"
	"go.fluxy.net/undeck/internal"
	"strconv"
	"strings"
)

const (
	UnknownSuit = iota
	Spade
	Diamond
	Club
	Heart
)

const (
	UnknownRank rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

func init() {
	var (
		_ undeck.Rank = rank(1)
		_ undeck.Suit = suit('S')
	)
}

// rank of a card 1-13
type rank int

func (r rank) String() string {
	switch {
	case r == 1:
		return "ACE"
	case r == 11:
		return "JACK"
	case r == 12:
		return "QUEEN"
	case r == 13:
		return "KING"
	case r > 0 && r < 11:
		return strconv.Itoa(int(r))
	}

	return "!"
}

func (r rank) Short() string {
	switch {
	case r == 1:
		return "A"
	case r == 11:
		return "J"
	case r == 12:
		return "Q"
	case r == 13:
		return "K"
	case r > 1 && r < 11:
		return strconv.Itoa(int(r))
	}

	return "!"
}

func (r rank) Validate() error {
	if r < 1 || r > 13 {
		return undeck.ErrInvalidRank
	}

	return nil
}

func rankFromString(s string) (rank, error) {
	switch s {
	case "A":
		return Ace, nil
	case "J":
		return Jack, nil
	case "Q":
		return Queen, nil
	case "K":
		return King, nil
	}

	var raw, err = strconv.Atoi(s)

	if err != nil {
		return 0, err
	}

	if raw >= 1 && raw <= 13 {
		return rank(raw), nil
	}

	return UnknownRank, undeck.ErrInvalidRank
}

// suit in a standard card deck - Spade, Diamond, Club and Heart
type suit rune

func (s suit) String() string {
	switch s {
	case Spade:
		return "SPADES"
	case Diamond:
		return "DIAMONDS"
	case Club:
		return "CLUBS"
	case Heart:
		return "HEARTS"
	}

	return "?"
}

func (s suit) Short() string {
	switch s {
	case Spade:
		return "S"
	case Diamond:
		return "D"
	case Club:
		return "C"
	case Heart:
		return "H"
	}

	return "?"
}

func (s suit) Validate() error {
	switch s {
	case Spade, Diamond, Club, Heart:
		return undeck.ErrInvalidSuit
	}

	return nil
}

func suitFromString(s string) (suit, error) {
	s = strings.ToUpper(s)

	switch s {
	case "S", "SPADE":
		return Spade, nil
	case "D", "DIAMOND":
		return Diamond, nil
	case "C", "CLUB":
		return Club, nil
	case "H", "HEART":
		return Heart, nil
	}

	return UnknownSuit, undeck.ErrInvalidSuit
}

// FromString returns a card from short hand string e.g. 2H will return 2 of Hearts
func FromString(s string) (undeck.Card, error) {
	var (
		err  error
		card undeck.Card

		head, suffix = internal.HeadSuffix(s)
	)

	card.Rank, err = rankFromString(head)
	if err != nil {
		return card, err
	}

	card.Suit, err = suitFromString(suffix)
	if err != nil {
		return card, err
	}

	return card, nil
}

// All returns the complete set of cards sequentially
func All() []undeck.Card {
	var (
		i     int
		cards []undeck.Card

		suits = []suit{
			Spade,
			Diamond,
			Club,
			Heart,
		}

		ranks = []rank{
			Ace,
			Two,
			Three,
			Four,
			Five,
			Six,
			Seven,
			Eight,
			Nine,
			Ten,
			Jack,
			Queen,
			King,
		}
	)

	for s := range suits {
		for r := range ranks {
			cards = append(cards, undeck.Card{
				Rank: ranks[r],
				Suit: suits[s],
			})

			i++
		}
	}

	return cards
}
