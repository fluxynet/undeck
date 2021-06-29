package undeck

import (
	"testing"
)

type testranksuit struct {
	str   string
	short string
	err   error
}

func (t testranksuit) String() string {
	return t.str
}

func (t testranksuit) Short() string {
	return t.short
}

func (t testranksuit) Validate() error {
	return t.err
}

func testcard(rankStr, rankShort, suitStr, suitShort string) Card {
	return Card{
		Rank: testranksuit{str: rankStr, short: rankShort},
		Suit: testranksuit{str: suitStr, short: suitShort},
	}
}

func assertCardSlicesEqual(t *testing.T, a, b []Card) bool {
	if len(a) != len(b) {
		t.Errorf("length a = %d, length b = %d\n", len(a), len(b))
		return false
	}

	for i := range a {
		if a[i].Rank.String() != b[i].Rank.String() {
			t.Errorf("rank not same, want %s, got %s", a[i].Rank.String(), b[i].Rank.String())
			return false
		}

		if a[i].Suit.String() != b[i].Suit.String() {
			t.Errorf("rank not same, want %s, got %s", a[i].Suit.String(), b[i].Suit.String())
			return false
		}
	}

	return true
}

func TestDeck_Add(t *testing.T) {
	type fields struct {
		ID         string
		IsShuffled bool
		Shuffler   ShufflerFunc
		cards      []Card
	}
	type args struct {
		cards []Card
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   Deck
	}{
		{
			name: "empty add none",
			fields: fields{
				ID:         "1",
				IsShuffled: true,
				Shuffler:   nil,
				cards:      []Card{},
			},
			args: args{cards: nil},
			want: Deck{
				ID:         "1",
				IsShuffled: true,
				Shuffler:   nil,
				cards:      []Card{},
			},
		},
		{
			name: "empty add one",
			fields: fields{
				ID:         "1",
				IsShuffled: true,
				Shuffler:   nil,
				cards:      []Card{},
			},
			args: args{cards: []Card{
				testcard("Ace", "A", "Hearts", "H"),
			}},
			want: Deck{
				ID:         "1",
				IsShuffled: true,
				Shuffler:   nil,
				cards: []Card{
					testcard("Ace", "A", "Hearts", "H"),
				},
			},
		},
		{
			name: "two add one",
			fields: fields{
				ID:         "1",
				IsShuffled: true,
				Shuffler:   nil,
				cards: []Card{
					testcard("Ace", "A", "Hearts", "H"),
					testcard("King", "K", "Hearts", "H"),
				},
			},
			args: args{
				cards: []Card{
					testcard("Queen", "Q", "Hearts", "H"),
				},
			},
			want: Deck{
				ID:         "1",
				IsShuffled: true,
				Shuffler:   nil,
				cards: []Card{
					testcard("Ace", "A", "Hearts", "H"),
					testcard("King", "K", "Hearts", "H"),
					testcard("Queen", "Q", "Hearts", "H"),
				},
			},
		},
		{
			name: "one add two",
			fields: fields{
				ID:         "1",
				IsShuffled: true,
				Shuffler:   nil,
				cards: []Card{
					testcard("Ace", "A", "Hearts", "H"),
				},
			},
			args: args{
				cards: []Card{
					testcard("King", "K", "Hearts", "H"),
					testcard("Queen", "Q", "Hearts", "H"),
				},
			},
			want: Deck{
				ID:         "1",
				IsShuffled: true,
				Shuffler:   nil,
				cards: []Card{
					testcard("Ace", "A", "Hearts", "H"),
					testcard("King", "K", "Hearts", "H"),
					testcard("Queen", "Q", "Hearts", "H"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Deck{
				ID:         tt.fields.ID,
				IsShuffled: tt.fields.IsShuffled,
				Shuffler:   tt.fields.Shuffler,
				cards:      tt.fields.cards,
			}

			got := d.Add(tt.args.cards...)

			if !assertCardSlicesEqual(t, tt.want.cards, got.cards) {
				t.Errorf("tt.want != got")
			}

			if !assertCardSlicesEqual(t, tt.fields.cards, d.cards) {
				t.Errorf("original deck changed: tt.fields.cards != d.cards")
			}
		})
	}
}
