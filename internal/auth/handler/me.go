package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sxd0/SweetTweet/internal/auth/middleware"
)

// MeHandler returns user info from token
//
// @Summary Get current user
// @Description Returns current user ID from JWT
// @Tags auth
// @Produce json
// @Success 200 {object} model.MeResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /me [get]
func MeHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey)
	if userID == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var uid int

	switch v := userID.(type) {
	case float64:
		uid = int(v)
	case int:
		uid = v
	case int64:
		uid = int(v)
	case json.Number:
		i, err := v.Int64()
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		uid = int(i)
	default:
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	resp := map[string]interface{}{
		"user_id": uid,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
