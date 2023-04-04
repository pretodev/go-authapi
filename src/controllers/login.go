package controllers

import (
	"encoding/json"
	"main/src/data"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"senha"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user data.User
	for _, u := range data.Users {
		if loginRequest.Email == u.Email && loginRequest.Password == u.Password() {
			user = u
			break
		}
	}

	if user.ID == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// sign token
	tokenString, err := user.CreateToken()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// return token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var loginResponse LoginResponse
	loginResponse.Token = tokenString
	err = json.NewEncoder(w).Encode(loginResponse)
	if err != nil {
		return
	}
}
