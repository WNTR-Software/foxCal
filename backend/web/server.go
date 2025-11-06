package web

import (
	"context"
	"fmt"
	"net/http"

	webutils "git.mstar.dev/mstar/goutils/http"
	"github.com/rs/zerolog"
)

//go:generate go tool swag init -generalInfo ./server.go -output swagger

// @title Foxcal API
// @version 0.1
// @description The API for the Foxcal service

// @license.name European Union Public License 1.2
// @license.url https://eupl.eu/1.2/en/

// @host foxcal.example.com
// @BasePath /

// Server is a wrapper struct to make a clean shutdown
// easier
type Server struct {
	server *http.Server
}

// Rfc9457Placeholder placeholder for problem details
type Rfc9457Placeholder struct {
	Type      string `json:"type"`
	Title     string `json:"title"`
	Status    int    `json:"status"`
	Reference string `json:"reference"`
}

func NewServer(logger zerolog.Logger) *Server {
	mux := http.NewServeMux()
	addRoutes(mux)

	finalHandler := webutils.ChainMiddlewares(
		mux,
		webutils.BuildLoggingMiddleware(&logger, true, []string{}, map[string]string{}),
	)
	return &Server{
		server: &http.Server{
			Handler: finalHandler,
		},
	}
}

func (s *Server) Run(addr string) error {
	s.server.Addr = addr
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Hello there")
	})

	mux.HandleFunc("GET /sampleKv/{id}", handleSampleKvGet)
	mux.HandleFunc("POST /sampleKv", handleSampleKvPost)
}
