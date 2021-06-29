package french

import (
	"go.fluxy.net/undeck"
	"go.fluxy.net/undeck/internal"
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
	switch r {
	case 1:
		return "ACE"
	case 2:
		return "TWO"
	case 3:
		return "THREE"
	case 4:
		return "FOUR"
	case 5:
		return "FIVE"
	case 6:
		return "SIX"
	case 7:
		return "SEVEN"
	case 8:
		return "EIGHT"
	case 9:
		return "NINE"
	case 10:
		return "TEN"
	case 11:
		return "JACK"
	case 12:
		return "QUEEN"
	case 13:
		return "KING"
	}

	return "!"
}

func (r rank) Short() string {
	switch r {
	case 1:
		return "A"
	case 2:
		return "2"
	case 3:
		return "3"
	case 4:
		return "4"
	case 5:
		return "5"
	case 6:
		return "6"
	case 7:
		return "7"
	case 8:
		return "8"
	case 9:
		return "9"
	case 10:
		return "T"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
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
	case "2":
		return Two, nil
	case "3":
		return Three, nil
	case "4":
		return Four, nil
	case "5":
		return Five, nil
	case "6":
		return Six, nil
	case "7":
		return Seven, nil
	case "8":
		return Eight, nil
	case "9":
		return Nine, nil
	case "T":
		return Ten, nil
	case "J":
		return Jack, nil
	case "Q":
		return Queen, nil
	case "K":
		return King, nil
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
