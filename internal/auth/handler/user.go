package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sxd0/SweetTweet/internal/auth/service"
)

type RegisterHandler struct {
	Service *service.RegisterService
}

func NewRegisterHandler(s *service.RegisterService) *RegisterHandler {
	return &RegisterHandler{Service: s}
}

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err = h.Service.Register(req.Email, req.Password)
	if err != nil {
		fmt.Println("REGISTRATION ERROR:", err)
		http.Error(w, "registration failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user registered"))
}
