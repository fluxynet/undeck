package undeck

// Deck contains cards
type Deck interface {
	// ID returns the id of the deck
	ID() string
	// IsShuffled returns the shuffle status of the deck
	IsShuffled() bool
	// Remaining cards count
	Remaining() int
	// Cards in the deck
	Cards() []Card

	// Add cards to the deck
	Add(cards ...Card) Deck
	// Draw cards from the deck
	Draw(count int) (Deck, []Card, error)
	// Shuffle returns a shuffled copy of the deck
	Shuffle() Deck
	// Duplicate returns a copy of the deck
	Duplicate() Deck
}
