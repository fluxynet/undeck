package internal

import (
	"go.fluxy.net/undeck"
	"testing"
)

type RepoDumper interface {
	Dump() map[string]undeck.Deck
}

func AssertReposEqual(t *testing.T, a, b undeck.Repo) bool {
	var x, y map[string]undeck.Deck

	if (a == nil) != (b == nil) {
		t.Errorf("one is nil and one is not\nwant= %v\ngot  = %v", a, b)
	}

	if r, ok := a.(RepoDumper); ok {
		x = r.Dump()
	} else {
		t.Fatal("repo does not implement RepoDumper; not suported")
	}

	if r, ok := b.(RepoDumper); ok {
		y = r.Dump()
	} else {
		t.Fatal("repo does not implement RepoDumper; not suported")
	}

	if len(x) != len(y) {
		t.Errorf("repo length different. want = %d, got = %d", len(x), len(y))
		return false
	}

	for id, dx := range x {
		var dy, ok = y[id]

		if !ok {
			t.Errorf("id %s present in want, not in got", id)
			return false
		}

		var equal = AssertDecksEqual(t, dx, dy)
		if !equal {
			t.Errorf("deck %s not same\nwant = %v\ngot  = %v", id, dx, dy)
			return false
		}
	}

	return true
}

func AssertCardSlicesEqual(t *testing.T, a, b []undeck.Card) bool {
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

func AssertDecksEqual(t *testing.T, a, b undeck.Deck) bool {
	if (a == nil) != (b == nil) {
		t.Errorf("one is nil and one is not\nwant= %v\ngot  = %v", a, b)
	}

	if a.ID() != b.ID() {
		t.Errorf("ids not same\nwant = %s\ngot  = %s", a.ID(), b.ID())
		return false
	}

	if a.IsShuffled() != b.IsShuffled() {
		t.Errorf("shuffled not same\nwant = %t\ngot  = %t", a.IsShuffled(), b.IsShuffled())
		return false
	}

	if a.Remaining() != b.Remaining() {
		t.Errorf("remaining not same\nwant = %d\ngot  = %d", a.Remaining(), b.Remaining())
		return false
	}

	return AssertCardSlicesEqual(t, a.Cards(), b.Cards())
}
