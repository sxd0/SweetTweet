package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sxd0/SweetTweet/internal/auth/service"
)

// RegisterHandler handles user registration
//
// @Summary Register a new user
// @Description Register a new user by email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param input body model.RegisterInput true "User credentials"
// @Success 200 {object} model.UserResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /register [post]
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
