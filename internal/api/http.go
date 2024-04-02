package api

import (
	"chatsearch/internal/adapters/http_handler"
	"chatsearch/internal/domain"
	"context"
	"fmt"
	"net/http"
	"time"
)

type HttpServer struct {
	server *http.Server
}

func NewHTTPServer(core *domain.SearchEngine) *HttpServer {
	handler := http_handler.NewHandler(core)
	server := &http.Server{
		Addr:    ":9999",
		Handler: handler,
	}

	return &HttpServer{server: server}
}

func (s *HttpServer) Run() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			fmt.Printf("%s\n", err)
		}
	}()
}

func (s *HttpServer) Stop() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	go func() {
		defer cancel()
		if err := s.server.Shutdown(ctx); err != nil {
			fmt.Printf("failed to stop http server: %s\n", err)
		}
	}()

	return ctx
}
