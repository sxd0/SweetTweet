package main

import (
	"fmt"
	"net/http"

	"github.com/sxd0/SweetTweet/pkg/config"
	"github.com/sxd0/SweetTweet/pkg/db"
	"github.com/sxd0/SweetTweet/pkg/logger"

	"github.com/sxd0/SweetTweet/internal/auth/handler"
	"github.com/sxd0/SweetTweet/internal/auth/middleware"
	"github.com/sxd0/SweetTweet/internal/auth/repository"
	"github.com/sxd0/SweetTweet/internal/auth/service"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/sxd0/SweetTweet/docs"
)

// @title Auth Service API
// @version 1.0
// @description SweetTweet
// @host localhost:8081
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

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

	loginService := service.NewLoginService(repo)
	loginHandler := handler.NewLoginHandler(loginService)
	mux.Handle("/login", loginHandler)
	
	mux.Handle("/debug-token", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Authorization: " + auth))
	}))
	mux.Handle("/me", middleware.JWTMiddleware(http.HandlerFunc(handler.MeHandler)))
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	log.Info("Auth service starting", "port", cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), mux)
	if err != nil {
		log.Error("Server failed to start", "error", err)
	}
}
