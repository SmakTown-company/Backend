package models

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		Handler:        handler,
	}
	return s.httpServer.ListenAndServe()
}
func (s *Server) Close(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
}
