package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"
	"go.fluxy.net/undeck/decks/decks"
	"go.fluxy.net/undeck/repo/memory"
	wchi "go.fluxy.net/undeck/web/chi"
	"go.fluxy.net/undeck/web/draw"
	"log"
	"net/http"
)

// Server over http
type Server struct {
	Port string
}

func (s Server) serveCmd(cmd *cobra.Command, args []string) {
	var (
		repo  = memory.New(decks.New)
		drawg = draw.New(repo, wchi.IDGetter)
		mux   = chi.NewMux()
	)

	mux.Route("/draw", func(r chi.Router) {
		r.Post("/deck", drawg.Create)
		r.Get("/deck/{id}", drawg.Open)
		r.Patch("/deck/{id}", drawg.Draw)
	})

	log.Println("Starting server on http://127.0.0.1:" + s.Port)
	http.ListenAndServe(":"+s.Port, mux)
}
