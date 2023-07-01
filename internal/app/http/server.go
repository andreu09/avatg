package http

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	engine *http.Server
}

type ServerConfig struct {
	Addr string
}

func ProvideServer(cfg ServerConfig, router http.Handler) (*Server, error) {
	httpServer := &http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	return &Server{engine: httpServer}, nil
}

func (s *Server) Run() error {
	return s.engine.ListenAndServe()
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.engine.Shutdown(ctx); err != nil {
		log.Println("server shutdown error:", err)
	}
	log.Println("stop server")
}
