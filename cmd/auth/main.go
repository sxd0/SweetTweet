package main

import (
	"fmt"
	"net/http"

	"github.com/sxd0/SweetTweet/pkg/config"
	"github.com/sxd0/SweetTweet/pkg/logger"
)

func main() {
	cfg := config.Load()
	log := logger.NewLogger()

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	log.Info("Auth service starting", "port", cfg.Port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), mux)
	if err != nil {
		log.Error("Server failed to start", "error", err)
	}
}
