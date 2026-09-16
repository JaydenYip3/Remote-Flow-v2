package server

import (
	"net/http"
	"time"

	"github.com/jaydenyip/remote_flow_v2/backend/internal/config"
)

type Server struct {
	http *http.Server
}

func New(cfg config.Config) *Server {
	s := Server{
		http: &http.Server{
			Addr:              cfg.Address,
			Handler:           routes(),
			ReadHeaderTimeout: 5 * time.Second,
		},
	}

	return &s
}

func (s *Server) Run() error {
	return s.http.ListenAndServe()
}
