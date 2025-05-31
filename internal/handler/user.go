package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nunusavi/task-manager/internal/service"
)

type RegisterRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
type RegisterResponse struct {
	Message string `json:"message"`
	UserID  int    `json:"user_id,omitempty"`
}
type LoginRequest struct{
	Email string `json:"email"`
	Password string `json:"password"`
}
type LoginResponse struct{
	Token string `json:"token"`
}
func RegisterUserHandler(w http.ResponseWriter, r *http.Request){
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := service.RegisterUser(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := RegisterResponse{
		Message: "User registered successfully",
		UserID:  user.ID,
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Login function 

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := service.LoginUser(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	resp := LoginResponse{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}