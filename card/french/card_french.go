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
		_ undeck.Card = card{}
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
		return "SPADE"
	case Diamond:
		return "DIAMOND"
	case Club:
		return "CLUB"
	case Heart:
		return "HEART"
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

// card consisting of a value and a suitFromString
type card struct {
	rank rank
	suit suit
}

func (c card) Rank() undeck.Rank {
	return c.rank
}

func (c card) Suit() undeck.Suit {
	return c.suit
}

func (c card) Duplicate() undeck.Card {
	return card{
		rank: c.rank,
		suit: c.suit,
	}
}

// FromString returns a card from short hand string e.g. 2H will return 2 of Hearts
func FromString(s string) (undeck.Card, error) {
	var (
		err  error
		card card

		head, suffix = internal.HeadSuffix(s)
	)

	card.rank, err = rankFromString(head)
	if err != nil {
		return nil, err
	}

	card.suit, err = suitFromString(suffix)
	if err != nil {
		return nil, err
	}

	return &card, nil
}

// All returns the complete set of cards sequentially
func All() []undeck.Card {
	var (
		i     int
		cards = make([]undeck.Card, 52)
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
			cards[i] = card{
				rank: ranks[r],
				suit: suits[s],
			}

			i++
		}
	}

	return cards
}
