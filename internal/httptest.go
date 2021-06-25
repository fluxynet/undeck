package internal

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func AssertHttp(t *testing.T, w *httptest.ResponseRecorder, wStatus int, wHeaders http.Header, wBody string) {
	if w.Code != wStatus {
		t.Errorf("wStatus not as wanted: want = %d, got = %d", wStatus, w.Code)
	}

	if w.Body.String() != wBody {
		var replacer = strings.NewReplacer(
			"\r", "[R]", "\n", "[N]",
		)

		t.Errorf(
			"wBody not same\nwant = %s\ngot  = %s",
			replacer.Replace(wBody),
			replacer.Replace(w.Body.String()),
		)
	}

	var headers = w.Header()

	if len(wHeaders) != len(headers) {
		t.Errorf("headers count: want = %d; %v\ngot  = %d; %v", len(wHeaders), wHeaders, len(headers), headers)

		return
	}

	for k, wv := range wHeaders {
		var gv = headers[k]

		if len(wv) != len(gv) {
			t.Errorf("wHeaders [%s] length: want = %d, got = %d", k, len(wv), len(gv))
			continue
		}

		for i := range wv {
			if wv[i] != gv[i] {
				t.Errorf("wHeaders[%s][%d]: want = %s, got = %s", k, i, wv[i], gv[i])
			}
		}
	}
}

type HttpTestRequest struct {
	Path   string
	Method string
	Header http.Header
	Body   string
}

type HttpTestWant struct {
	Status int
	Header http.Header
	Body   string
}

// HttpTest case with assertions
type HttpTest struct {
	Name string

	// Handler is the function we are testing
	Handler http.HandlerFunc

	// Request details we will be sending
	Request HttpTestRequest

	// Want stuffs
	Want HttpTestWant
}

// Assert all as per Want
func (h HttpTest) Assert(t *testing.T) {
	var b io.Reader

	if len(h.Request.Body) != 0 {
		b = strings.NewReader(h.Request.Body)
	}

	r, err := http.NewRequest(h.Request.Method, h.Request.Path, b)
	if err != nil {
		t.Errorf("[%s] failed to init request: %s", h.Name, err.Error())
		return
	}

	w := httptest.NewRecorder()

	h.Handler.ServeHTTP(w, r)

	var (
		result  = w.Result()
		status  = result.StatusCode
		headers = w.Header()

		lg = len(headers)
		lw = len(h.Want.Header)
		wb = h.Want.Body
		gb = w.Body.String()
	)

	if status != h.Want.Status {
		t.Errorf("[%s] status not as wanted: want = %d, got = %d", h.Name, h.Want.Status, status)
	}

	if wb != gb {
		var replacer = strings.NewReplacer(
			"\r", "[R]", "\n", "[N]",
		)

		wb = replacer.Replace(wb)
		gb = replacer.Replace(gb)

		t.Errorf("[%s] body not as wanted.\nwant = %s\ngot  = %s", h.Name, wb, gb)
	}

	if lw != lg {
		t.Errorf("[%s] headers count: want = %d, got = %d", h.Name, lw, lg)
		return
	}

	for k, wv := range h.Want.Header {
		var gv = headers[k]

		if len(wv) != len(gv) {
			t.Errorf("[%s] header [%s] length: want = %d, got = %d", h.Name, k, len(wv), len(gv))
			continue
		}

		for i := range wv {
			if wv[i] != gv[i] {
				t.Errorf("[%s] header[%s][%d]: want = %s, got = %s", h.Name, k, i, wv[i], gv[i])
			}
		}
	}

	return
}
