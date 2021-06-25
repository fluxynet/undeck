package draw

import (
	"go.fluxy.net/undeck"
	"go.fluxy.net/undeck/cards"
	"go.fluxy.net/undeck/cards/french"
	"go.fluxy.net/undeck/web"
	"net/http"
	"strconv"
)

func New(repo undeck.Repo, idGetter web.IDGetter) *Draw {
	return &Draw{
		repo:     repo,
		idGetter: idGetter,
	}
}

// Draw game served via web consists of drawing cards from a 52 card french deck
type Draw struct {
	repo     undeck.Repo
	idGetter web.IDGetter
}

type createResponse struct {
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

func (s *Draw) Create(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = r.Context()
		query = r.URL.Query()

		res createResponse

		deck, err = s.repo.Create(ctx)
	)

	if err != nil {
		web.JsonError(w, http.StatusInternalServerError, err)
		return
	}

	if rawCards := query.Get("cards"); rawCards == "" {
		deck = deck.Add(french.All()...)
	} else if cs, err := cards.FromString(french.FromString, rawCards); err == nil {
		deck = deck.Add(cs...)
	} else {
		web.JsonError(w, http.StatusBadRequest, err)
		return
	}

	// shuffle it
	if rawShuffle := query.Get("shuffle"); rawShuffle == "" {
		// ignore it
	} else if b, err := strconv.ParseBool(rawShuffle); err != nil {
		web.JsonError(w, http.StatusBadRequest, err)
		return
	} else if b {
		deck = deck.Shuffle()
	}

	deck, err = s.repo.Save(ctx, deck)

	if err != nil {
		web.JsonError(w, http.StatusInternalServerError, err)
		return
	}

	res = createResponse{
		DeckID:    deck.ID(),
		Shuffled:  deck.IsShuffled(),
		Remaining: deck.Remaining(),
	}

	web.Json(w, res)
}

type openResponse struct {
	DeckID    string             `json:"deck_id"`
	Shuffled  bool               `json:"shuffled"`
	Remaining int                `json:"remaining"`
	Cards     []undeck.CardState `json:"cards"`
}

func (s *Draw) Open(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		deck undeck.Deck
		res  openResponse

		id, err = s.idGetter(r)
	)

	if err != nil {
		web.JsonError(w, http.StatusBadRequest, err)
		return
	}

	if deck, err = s.repo.Find(ctx, id); err == undeck.ErrDeckNotFound {
		web.JsonError(w, http.StatusNotFound, err)
		return
	} else if err != nil {
		web.JsonError(w, http.StatusInternalServerError, err)
		return
	}

	res = openResponse{
		DeckID:    deck.ID(),
		Shuffled:  deck.IsShuffled(),
		Remaining: deck.Remaining(),
	}

	for _, c := range deck.Cards() {
		res.Cards = append(res.Cards, undeck.ToCardState(c))
	}

	web.Json(w, res)
}

type drawResponse struct {
	Cards []undeck.CardState `json:"cards"`
}

func (s *Draw) Draw(w http.ResponseWriter, r *http.Request) {
	var (
		count    int
		cardlist []undeck.Card
		res      drawResponse
		deck     undeck.Deck

		ctx     = r.Context()
		id, err = s.idGetter(r)
	)

	if err != nil {
		web.JsonError(w, http.StatusBadRequest, err)
		return
	}

	if raw := r.URL.Query().Get("count"); raw == "" {
		count = 1
	} else if count, err = strconv.Atoi(raw); err != nil {
		web.JsonError(w, http.StatusBadRequest, err)
		return
	}

	if d, err := s.repo.Find(ctx, id); err == nil {
		deck = d
	} else if err == undeck.ErrDeckNotFound {
		web.JsonError(w, http.StatusNotFound, err)
		return
	} else {
		web.JsonError(w, http.StatusInternalServerError, err)
		return
	}

	deck, cardlist, err = deck.Draw(count)
	if err == undeck.ErrNotEnoughCards {
		web.JsonError(w, http.StatusBadRequest, err)
		return
	} else if err != nil {
		web.JsonError(w, http.StatusInternalServerError, err)
		return
	}

	_, err = s.repo.Save(ctx, deck)
	if err != nil {
		web.JsonError(w, http.StatusInternalServerError, err)
		return
	}

	for i := range cardlist {
		res.Cards = append(res.Cards, undeck.ToCardState(cardlist[i]))
	}

	web.Json(w, res)
}
