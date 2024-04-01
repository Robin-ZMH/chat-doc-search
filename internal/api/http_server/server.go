package http_server

import (
	"chatsearch/internal/domain"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	server *http.Server
}

func New(core *domain.SearchEngine) *HttpServer {
	handler := NewHandler(core)

	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/engine/query", handler.Query)
	router.POST("/engine/query", handler.Query)
	router.POST("/engine/insert", handler.Insert)
	router.PUT("/engine", handler.Update)
	router.PATCH("/engine", handler.Update)
	router.DELETE("/engine", handler.Delete)

	server := &http.Server{
		Addr:    ":9999",
		Handler: router,
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
