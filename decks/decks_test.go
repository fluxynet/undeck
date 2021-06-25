package decks

import (
	"go.fluxy.net/undeck"
	"go.fluxy.net/undeck/cards/french"
	"testing"
)

func assertCardSlicesEqual(t *testing.T, a, b []undeck.Card) bool {
	if len(a) != len(b) {
		t.Errorf("length a = %d, length b = %d\n", len(a), len(b))
		return false
	}

	for i := range a {
		if a[i].Rank() != b[i].Rank() {
			t.Errorf("rank [%d] not same, want %d, got %d", i, a[i].Rank(), b[i].Rank())
			return false
		}

		if a[i].Suit() != b[i].Suit() {
			t.Errorf("rank [%d] not same, want %s, got %s", i, a[i].Suit().String(), b[i].Suit().String())
			return false
		}
	}

	return true
}

func TestAddCards(t *testing.T) {
	type args struct {
		cards []undeck.Card
		more  []undeck.Card
	}

	var tests = []struct {
		name string
		args args
		want []undeck.Card
	}{
		{
			name: "empty add none",
			args: args{
				cards: []undeck.Card{},
				more:  []undeck.Card{},
			},
			want: []undeck.Card{},
		},
		{
			name: "empty add one",
			args: args{
				cards: []undeck.Card{},
				more: []undeck.Card{
					french.MustString("AH"),
				},
			},
			want: []undeck.Card{
				french.MustString("AH"),
			},
		},
		{
			name: "two add one",
			args: args{
				cards: []undeck.Card{
					french.MustString("AH"),
					french.MustString("KH"),
				},
				more: []undeck.Card{
					french.MustString("QH"),
				},
			},
			want: []undeck.Card{
				french.MustString("AH"),
				french.MustString("KH"),
				french.MustString("QH"),
			},
		},
		{
			name: "one add two",
			args: args{
				cards: []undeck.Card{
					french.MustString("AH"),
				},
				more: []undeck.Card{
					french.MustString("KH"),
					french.MustString("QH"),
				},
			},
			want: []undeck.Card{
				french.MustString("AH"),
				french.MustString("KH"),
				french.MustString("QH"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got = AddCards(tt.args.cards, tt.args.more...)

			if !assertCardSlicesEqual(t, tt.want, got) {
				t.Errorf("tt.want != got")
			}

		})
	}
}

func TestRemoveCards(t *testing.T) {
	type args struct {
		cards []undeck.Card
		count int
	}

	tests := []struct {
		name          string
		args          args
		wantRemaining []undeck.Card
		wantRemoved   []undeck.Card
		wantErr       error
	}{
		{
			name: "empty remove none",
			args: args{
				cards: []undeck.Card{},
				count: 0,
			},
			wantRemaining: []undeck.Card{},
			wantRemoved:   []undeck.Card{},
			wantErr:       nil,
		},
		{
			name: "empty remove one",
			args: args{
				cards: []undeck.Card{},
				count: 1,
			},
			wantRemaining: []undeck.Card{},
			wantRemoved:   []undeck.Card{},
			wantErr:       undeck.ErrNotEnoughCards,
		},
		{
			name: "one remove two",
			args: args{
				cards: []undeck.Card{
					french.MustString("KH"),
				},
				count: 2,
			},
			wantRemaining: []undeck.Card{
				french.MustString("KH"),
			},
			wantRemoved: []undeck.Card{},
			wantErr:     undeck.ErrNotEnoughCards,
		},
		{
			name: "two remove two",
			args: args{
				cards: []undeck.Card{
					french.MustString("QH"),
					french.MustString("KH"),
				},
				count: 2,
			},
			wantRemaining: []undeck.Card{},
			wantRemoved: []undeck.Card{
				french.MustString("QH"),
				french.MustString("KH"),
			},
			wantErr: nil,
		},
		{
			name: "two remove one",
			args: args{
				cards: []undeck.Card{
					french.MustString("QH"),
					french.MustString("KH"),
				},
				count: 1,
			},
			wantRemaining: []undeck.Card{
				french.MustString("KH"),
			},
			wantRemoved: []undeck.Card{
				french.MustString("QH"),
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotRemaining, gotRemoved, err = RemoveCards(tt.args.cards, tt.args.count)

			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("RemoveCards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !assertCardSlicesEqual(t, gotRemaining, tt.wantRemaining) {
				t.Errorf("gotRemaining != tt.wantRemaining")
				return
			}

			if !assertCardSlicesEqual(t, gotRemoved, tt.wantRemoved) {
				t.Errorf("gotRemoved != tt.wantRemoved")
				return
			}
		})
	}
}
