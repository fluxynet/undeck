package chi

import (
	chi "github.com/go-chi/chi/v5"
	"go.fluxy.net/undeck/web"
	"net/http"
)

func IDGetter(r *http.Request) (string, error) {
	var id = chi.URLParam(r, "id")

	if id == "" {
		return "", web.ErrIDMissing
	}

	return id, nil
}
