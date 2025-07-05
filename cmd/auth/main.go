package main

import (
	"fmt"
	"net/http"

	"Go-pet/internal/pkg/config"
	"Go-pet/internal/pkg/logger"
)

func main() {
	cfg := config.Load()
	log := logger.NewLogger()
	defer log.Sync()

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	log.Info("Auth service starting", zap.String("port", cfg.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), mux)
	if err != nil {
		log.Fatal("Server failed to start", zap.Error(err))
	}
}
