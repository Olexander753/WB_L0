package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	htppServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	log.Println("Run server")
	s.htppServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 28,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.htppServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.htppServer.Shutdown(ctx)
}
