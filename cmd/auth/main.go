package main

import (
	"fmt"
	"net/http"

	"github.com/sxd0/SweetTweet/pkg/config"
	"github.com/sxd0/SweetTweet/pkg/logger"
	"github.com/sxd0/SweetTweet/pkg/db"

	"github.com/sxd0/SweetTweet/internal/auth/handler"
	"github.com/sxd0/SweetTweet/internal/auth/repository"
	"github.com/sxd0/SweetTweet/internal/auth/service"
)

func main() {
	cfg := config.Load()
	log := logger.NewLogger()

	database, err := db.Connect()
	if err != nil {
		log.Error("DB connection failed", "error", err)
		return
	}
	log.Info("Connected to DB")

	repo := repository.NewUserRepository(database)
	svc := service.NewRegisterService(repo)
	registerHandler := handler.NewRegisterHandler(svc)

	mux := http.NewServeMux()
	mux.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}))
	mux.Handle("/register", registerHandler)

	log.Info("Auth service starting", "port", cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), mux)
	if err != nil {
		log.Error("Server failed to start", "error", err)
	}
}

