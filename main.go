package main

import (
	"chatsearch/internal/adapters/mongodb"
	"chatsearch/internal/api"
	"chatsearch/internal/domain"
	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	repo := mongodb.NewMongoRepo("mongodb://mongodb:27017")
	core := domain.NewSearchEngine(repo)
	server := api.NewHTTPServer(core)
	server.Run()

	<-ctx.Done()
	stop()

	<-server.Stop().Done()
}
