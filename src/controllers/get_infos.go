package controllers

import (
	"encoding/json"
	"main/src/data"
	"net/http"
)

func GetInfos(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" || len(tokenString) < 7 {
		http.Error(w, "Token de autenticação não fornecido", http.StatusUnauthorized)
		return
	}

	tokenString = tokenString[7:]

	user, err := data.GetUserFromToken(tokenString)
	if err != nil {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}

	user.GenerateBalance()

	// Retornar as informações do usuário
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}
}
