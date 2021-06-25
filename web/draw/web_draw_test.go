package draw

import (
	"go.fluxy.net/undeck"
	"go.fluxy.net/undeck/cards"
	"go.fluxy.net/undeck/cards/french"
	"go.fluxy.net/undeck/decks/dummy"
	"go.fluxy.net/undeck/internal"
	"go.fluxy.net/undeck/repo"
	"go.fluxy.net/undeck/repo/memory"
	"go.fluxy.net/undeck/web"
	"net/http"
	"testing"
)

type fields struct {
	repo     undeck.Repo
	idGetter web.IDGetter
}

type test struct {
	fields fields
	after  undeck.Repo
	http   internal.HttpTest
}

func TestDraw_Create(t *testing.T) {
	var fakeFrenchShuffled = french.All()

	fakeFrenchShuffled[0], fakeFrenchShuffled[1] = fakeFrenchShuffled[1], fakeFrenchShuffled[0]

	var tests = []test{
		{
			fields: fields{
				repo: memory.NewWith(
					dummy.NewDecker(),
					repo.Sequential("1"),
				),
				idGetter: nil,
			},
			after: memory.NewWith(
				nil, nil, dummy.NewWithCards("1", false, french.All()...),
			),
			http: internal.HttpTest{
				Name:    "no arguments",
				Handler: nil,
				Request: internal.HttpTestRequest{
					Path:   "",
					Method: http.MethodPost,
					Header: http.Header{},
					Body:   "",
				},
				Want: internal.HttpTestWant{
					Status: http.StatusOK,
					Header: http.Header{
						"Content-Type": {web.ContentTypeJSON},
					},
					Body: `{"deck_id":"1","shuffled":false,"remaining":52}`,
				},
			},
		},
		{
			fields: fields{
				repo: memory.NewWith(
					dummy.NewDecker(),
					repo.Sequential("1"),
				),
				idGetter: nil,
			},
			after: memory.NewWith(nil, nil),
			http: internal.HttpTest{
				Name:    "bad card list param",
				Handler: nil,
				Request: internal.HttpTestRequest{
					Path:   "?cards=2Spade,3D,4C,5H",
					Method: http.MethodPost,
					Header: http.Header{},
					Body:   "",
				},
				Want: internal.HttpTestWant{
					Status: http.StatusBadRequest,
					Header: http.Header{
						"Content-Type": {web.ContentTypeJSON},
					},
					Body: `{"error":"strconv.Atoi: parsing \"2Spad\": invalid syntax"}`,
				},
			},
		},
		{
			fields: fields{
				repo: memory.NewWith(
					dummy.NewDecker(),
					repo.Sequential("1"),
				),
				idGetter: nil,
			},
			after: memory.NewWith(
				nil, nil, dummy.NewWithCards("1", false, cards.MustString(french.FromString, "2S,3D,4C,5H")...),
			),
			http: internal.HttpTest{
				Name:    "card list only",
				Handler: nil,
				Request: internal.HttpTestRequest{
					Path:   "?cards=2S,3D,4C,5H",
					Method: http.MethodPost,
					Header: http.Header{},
					Body:   "",
				},
				Want: internal.HttpTestWant{
					Status: http.StatusOK,
					Header: http.Header{
						"Content-Type": {web.ContentTypeJSON},
					},
					Body: `{"deck_id":"1","shuffled":false,"remaining":4}`,
				},
			},
		},
		{
			fields: fields{
				repo: memory.NewWith(
					dummy.NewDecker(),
					repo.Sequential("1"),
				),
				idGetter: nil,
			},
			after: memory.NewWith(nil, nil),
			http: internal.HttpTest{
				Name:    "bad shuffle param",
				Handler: nil,
				Request: internal.HttpTestRequest{
					Path:   "?shuffle=Yes",
					Method: http.MethodPost,
					Header: http.Header{},
					Body:   "",
				},
				Want: internal.HttpTestWant{
					Status: http.StatusBadRequest,
					Header: http.Header{
						"Content-Type": {web.ContentTypeJSON},
					},
					Body: `{"error":"strconv.ParseBool: parsing \"Yes\": invalid syntax"}`,
				},
			},
		},
		{
			fields: fields{
				repo: memory.NewWith(
					dummy.NewDecker(),
					repo.Sequential("1"),
				),
				idGetter: nil,
			},
			after: memory.NewWith(
				nil, nil, dummy.NewWithCards("1", true, fakeFrenchShuffled...),
			),
			http: internal.HttpTest{
				Name:    "shuffle only",
				Handler: nil,
				Request: internal.HttpTestRequest{
					Path:   "?shuffle=true",
					Method: http.MethodPost,
					Header: http.Header{},
					Body:   "",
				},
				Want: internal.HttpTestWant{
					Status: http.StatusOK,
					Header: http.Header{
						"Content-Type": {web.ContentTypeJSON},
					},
					Body: `{"deck_id":"1","shuffled":true,"remaining":52}`,
				},
			},
		},
		{
			fields: fields{
				repo: memory.NewWith(
					dummy.NewDecker(),
					repo.Sequential("1"),
				),
				idGetter: nil,
			},
			after: memory.NewWith(
				nil, nil, dummy.NewWithCards("1", true, cards.MustString(french.FromString, "9D,10S,8C,7H")...),
			),
			http: internal.HttpTest{
				Name:    "card list and shuffle",
				Handler: nil,
				Request: internal.HttpTestRequest{
					Path:   "?cards=10S,9D,8C,7H&shuffle=true",
					Method: http.MethodPost,
					Header: http.Header{},
					Body:   "",
				},
				Want: internal.HttpTestWant{
					Status: http.StatusOK,
					Header: http.Header{
						"Content-Type": {web.ContentTypeJSON},
					},
					Body: `{"deck_id":"1","shuffled":true,"remaining":4}`,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.http.Name, func(t *testing.T) {
			s := &Draw{
				repo:     tt.fields.repo,
				idGetter: tt.fields.idGetter,
			}

			tt.http.Handler = s.Create

			tt.http.Assert(t)
			internal.AssertReposEqual(t, tt.fields.repo, tt.after)
		})
	}
}

func TestDraw_Open(t *testing.T) {
	var tests = []test{
		{
			fields: fields{
				repo: memory.NewWith(
					nil,
					nil,
					dummy.NewWithCards("1", true, cards.MustString(french.FromString, "AH,JH,QH,KH")...),
				),
				idGetter: web.StaticIDGetter("", web.ErrIDMissing),
			},
			after: memory.NewWith(
				nil,
				nil,
				dummy.NewWithCards("1", true, cards.MustString(french.FromString, "AH,JH,QH,KH")...),
			),
			http: internal.HttpTest{
				Name:    "id missing",
				Handler: nil,
				Request: internal.HttpTestRequest{
					Path:   "",
					Method: "",
					Header: http.Header{},
					Body:   "",
				},
				Want: internal.HttpTestWant{
					Status: http.StatusBadRequest,
					Header: http.Header{
						"Content-Type": {web.ContentTypeJSON},
					},
					Body: `{"error":"id missing from request"}`,
				},
			},
		},
		{
			fields: fields{
				repo: memory.NewWith(
					nil,
					nil,
					dummy.NewWithCards("1", true, cards.MustString(french.FromString, "AH,JH,QH,KH")...),
				),
				idGetter: web.StaticIDGetter("2", nil),
			},
			after: memory.NewWith(
				nil,
				nil,
				dummy.NewWithCards("1", true, cards.MustString(french.FromString, "AH,JH,QH,KH")...),
			),
			http: internal.HttpTest{
				Name:    "non-existent deck",
				Handler: nil,
				Request: internal.HttpTestRequest{
					Path:   "",
					Method: "",
					Header: http.Header{},
					Body:   "",
				},
				Want: internal.HttpTestWant{
					Status: http.StatusNotFound,
					Header: http.Header{
						"Content-Type": {web.ContentTypeJSON},
					},
					Body: `{"error":"deck not found"}`,
				},
			},
		},
		{
			fields: fields{
				repo: memory.NewWith(
					nil,
					nil,
					dummy.NewWithCards("1", false),
				),
				idGetter: web.StaticIDGetter("1", nil),
			},
			after: memory.NewWith(
				nil,
				nil,
				dummy.NewWithCards("1", false),
			),
			http: internal.HttpTest{
				Name:    "empty deck",
				Handler: nil,
				Request: internal.HttpTestRequest{
					Path:   "",
					Method: "",
					Header: http.Header{},
					Body:   "",
				},
				Want: internal.HttpTestWant{
					Status: http.StatusOK,
					Header: http.Header{
						"Content-Type": {web.ContentTypeJSON},
					},
					Body: `{"deck_id":"1","shuffled":false,"remaining":0,"cards":null}`,
				},
			},
		},
		{
			fields: fields{
				repo: memory.NewWith(
					nil,
					nil,
					dummy.NewWithCards("1", true, cards.MustString(french.FromString, "AH,JH,QH,KH")...),
				),
				idGetter: web.StaticIDGetter("1", nil),
			},
			after: memory.NewWith(
				nil,
				nil,
				dummy.NewWithCards("1", true, cards.MustString(french.FromString, "AH,JH,QH,KH")...),
			),
			http: internal.HttpTest{
				Name:    "existing shuffled deck",
				Handler: nil,
				Request: internal.HttpTestRequest{
					Path:   "",
					Method: "",
					Header: http.Header{},
					Body:   "",
				},
				Want: internal.HttpTestWant{
					Status: http.StatusOK,
					Header: http.Header{
						"Content-Type": {web.ContentTypeJSON},
					},
					Body: `{"deck_id":"1","shuffled":true,"remaining":4,"cards":[{"value":"ACE","suit":"HEARTS","code":"AH"},{"value":"JACK","suit":"HEARTS","code":"JH"},{"value":"QUEEN","suit":"HEARTS","code":"QH"},{"value":"KING","suit":"HEARTS","code":"KH"}]}`,
				},
			},
		},
		{
			fields: fields{
				repo: memory.NewWith(
					nil,
					nil,
					dummy.NewWithCards("1", false, cards.MustString(french.FromString, "AH,JH,QH,KH")...),
				),
				idGetter: web.StaticIDGetter("1", nil),
			},
			after: memory.NewWith(
				nil,
				nil,
				dummy.NewWithCards("1", false, cards.MustString(french.FromString, "AH,JH,QH,KH")...),
			),
			http: internal.HttpTest{
				Name:    "existing non-shuffled deck",
				Handler: nil,
				Request: internal.HttpTestRequest{
					Path:   "",
					Method: "",
					Header: http.Header{},
					Body:   "",
				},
				Want: internal.HttpTestWant{
					Status: http.StatusOK,
					Header: http.Header{
						"Content-Type": {web.ContentTypeJSON},
					},
					Body: `{"deck_id":"1","shuffled":false,"remaining":4,"cards":[{"value":"ACE","suit":"HEARTS","code":"AH"},{"value":"JACK","suit":"HEARTS","code":"JH"},{"value":"QUEEN","suit":"HEARTS","code":"QH"},{"value":"KING","suit":"HEARTS","code":"KH"}]}`,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.http.Name, func(t *testing.T) {
			s := &Draw{
				repo:     tt.fields.repo,
				idGetter: tt.fields.idGetter,
			}

			tt.http.Handler = s.Open

			tt.http.Assert(t)
			internal.AssertReposEqual(t, tt.fields.repo, tt.after)
		})
	}
}

func TestDraw_Draw(t *testing.T) {
	var tests = []test{
		{
			fields: fields{
				repo: memory.NewWith(
					nil,
					nil,
					dummy.NewWithCards("1", true, cards.MustString(french.FromString, "AH,JH,QH,KH")...),
				),
				idGetter: web.StaticIDGetter("", web.ErrIDMissing),
			},
			after: memory.NewWith(
				nil,
				nil,
				dummy.NewWithCards("1", true, cards.MustString(french.FromString, "AH,JH,QH,KH")...),
			),
			http: internal.HttpTest{
				Name:    "id missing",
				Handler: nil,
				Request: internal.HttpTestRequest{
					Path:   "",
					Method: "",
					Header: http.Header{},
					Body:   "",
				},
				Want: internal.HttpTestWant{
					Status: http.StatusBadRequest,
					Header: http.Header{
						"Content-Type": {web.ContentTypeJSON},
					},
					Body: `{"error":"id missing from request"}`,
				},
			},
		},
		{
			fields: fields{
				repo: memory.NewWith(
					nil,
					nil,
					dummy.NewWithCards("1", true, cards.MustString(french.FromString, "AH,JH,QH,KH")...),
				),
				idGetter: web.StaticIDGetter("2", nil),
			},
			after: memory.NewWith(
				nil,
				nil,
				dummy.NewWithCards("1", true, cards.MustString(french.FromString, "AH,JH,QH,KH")...),
			),
			http: internal.HttpTest{
				Name:    "non-existent deck",
				Handler: nil,
				Request: internal.HttpTestRequest{
					Path:   "",
					Method: "",
					Header: http.Header{},
					Body:   "",
				},
				Want: internal.HttpTestWant{
					Status: http.StatusNotFound,
					Header: http.Header{
						"Content-Type": {web.ContentTypeJSON},
					},
					Body: `{"error":"deck not found"}`,
				},
			},
		},
		{
			fields: fields{
				repo: memory.NewWith(
					nil,
					nil,
					dummy.NewWithCards("1", true, cards.MustString(french.FromString, "AH,JH,QH,KH")...),
				),
				idGetter: web.StaticIDGetter("1", nil),
			},
			after: memory.NewWith(
				nil,
				nil,
				dummy.NewWithCards("1", true, cards.MustString(french.FromString, "AH,JH,QH,KH")...),
			),
			http: internal.HttpTest{
				Name:    "bad count param",
				Handler: nil,
				Request: internal.HttpTestRequest{
					Path:   "?count=ten",
					Method: "",
					Header: http.Header{},
					Body:   "",
				},
				Want: internal.HttpTestWant{
					Status: http.StatusBadRequest,
					Header: http.Header{
						"Content-Type": {web.ContentTypeJSON},
					},
					Body: `{"error":"strconv.Atoi: parsing \"ten\": invalid syntax"}`,
				},
			},
		},
		{
			fields: fields{
				repo: memory.NewWith(
					nil,
					nil,
					dummy.NewWithCards("1", false),
				),
				idGetter: web.StaticIDGetter("1", nil),
			},
			after: memory.NewWith(
				nil,
				nil,
				dummy.NewWithCards("1", false),
			),
			http: internal.HttpTest{
				Name:    "empty deck",
				Handler: nil,
				Request: internal.HttpTestRequest{
					Path:   "",
					Method: "",
					Header: http.Header{},
					Body:   "",
				},
				Want: internal.HttpTestWant{
					Status: http.StatusBadRequest,
					Header: http.Header{
						"Content-Type": {web.ContentTypeJSON},
					},
					Body: `{"error":"deck does not contain enough cards"}`,
				},
			},
		},
		{
			fields: fields{
				repo: memory.NewWith(
					nil,
					nil,
					dummy.NewWithCards("1", true, cards.MustString(french.FromString, "AH,JH,QH,KH")...),
				),
				idGetter: web.StaticIDGetter("1", nil),
			},
			after: memory.NewWith(
				nil,
				nil,
				dummy.NewWithCards("1", true, cards.MustString(french.FromString, "JH,QH,KH")...),
			),
			http: internal.HttpTest{
				Name:    "non-empty deck draw",
				Handler: nil,
				Request: internal.HttpTestRequest{
					Path:   "",
					Method: "",
					Header: http.Header{},
					Body:   "",
				},
				Want: internal.HttpTestWant{
					Status: http.StatusOK,
					Header: http.Header{
						"Content-Type": {web.ContentTypeJSON},
					},
					Body: `{"cards":[{"value":"ACE","suit":"HEARTS","code":"AH"}]}`,
				},
			},
		},
		{
			fields: fields{
				repo: memory.NewWith(
					nil,
					nil,
					dummy.NewWithCards("1", true, cards.MustString(french.FromString, "AH,JH,QH,KH")...),
				),
				idGetter: web.StaticIDGetter("1", nil),
			},
			after: memory.NewWith(
				nil,
				nil,
				dummy.NewWithCards("1", true, cards.MustString(french.FromString, "QH,KH")...),
			),
			http: internal.HttpTest{
				Name:    "non-empty deck draw 2",
				Handler: nil,
				Request: internal.HttpTestRequest{
					Path:   "?count=2",
					Method: "",
					Header: http.Header{},
					Body:   "",
				},
				Want: internal.HttpTestWant{
					Status: http.StatusOK,
					Header: http.Header{
						"Content-Type": {web.ContentTypeJSON},
					},
					Body: `{"cards":[{"value":"ACE","suit":"HEARTS","code":"AH"},{"value":"JACK","suit":"HEARTS","code":"JH"}]}`,
				},
			},
		},
		{
			fields: fields{
				repo: memory.NewWith(
					nil,
					nil,
					dummy.NewWithCards("1", true, cards.MustString(french.FromString, "AH,JH,QH,KH")...),
				),
				idGetter: web.StaticIDGetter("1", nil),
			},
			after: memory.NewWith(
				nil,
				nil,
				dummy.NewWithCards("1", true, cards.MustString(french.FromString, "AH,JH,QH,KH")...),
			),
			http: internal.HttpTest{
				Name:    "non-empty deck overdraw",
				Handler: nil,
				Request: internal.HttpTestRequest{
					Path:   "?count=100",
					Method: "",
					Header: http.Header{},
					Body:   "",
				},
				Want: internal.HttpTestWant{
					Status: http.StatusBadRequest,
					Header: http.Header{
						"Content-Type": {web.ContentTypeJSON},
					},
					Body: `{"error":"deck does not contain enough cards"}`,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.http.Name, func(t *testing.T) {
			s := &Draw{
				repo:     tt.fields.repo,
				idGetter: tt.fields.idGetter,
			}

			tt.http.Handler = s.Draw

			tt.http.Assert(t)
			internal.AssertReposEqual(t, tt.fields.repo, tt.after)
		})
	}
}
