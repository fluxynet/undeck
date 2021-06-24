package undeck

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidRank indicates that the card does not have a valid rank
	ErrInvalidRank = errors.New("card does not have a valid rank")

	// ErrInvalidSuit indicates that the card does not have a valid suit
	ErrInvalidSuit = errors.New("card does not have a valid suit")

	// ErrNotEnoughCards indicates that operations on a deck failed because the deck does not contain enough cards
	ErrNotEnoughCards = errors.New("deck does not contain enough cards")
)

// Rank of a card depending on the game being played, in a 52 french deck: Ace, 2-10, Jack, Queen and King
type Rank interface {
	fmt.Stringer
	Short() string
	Validate() error
}

// Suit of a card depending on the game being played, in a 52 french deck: Spade, Diamond, Club, Heart
type Suit interface {
	fmt.Stringer
	Short() string
	Validate() error
}

// Card used in games
type Card interface {
	Rank() Rank
	Suit() Suit
	Duplicate() Card
}

// CardState the state of a card, can be used for serialization
type CardState struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

// ToCardState returns the CardState representation of a card
func ToCardState(c Card) CardState {
	var (
		s CardState

		rank = c.Rank()
		suit = c.Suit()
	)

	s.Code = rank.Short() + suit.Short()
	s.Suit = suit.String()
	s.Value = rank.String()

	return s
}
