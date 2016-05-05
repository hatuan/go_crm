package controllers

import (
	"encoding/json"
	"erpvietnam/crm/auth"
	"erpvietnam/crm/models"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestLogin := new(models.LoginDTO)
	json.NewDecoder(r.Body).Decode(&requestLogin)

	responseStatus, token := auth.Login(requestLogin)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := new(models.LoginDTO)
	json.NewDecoder(r.Body).Decode(&requestUser)

	w.Header().Set("Content-Type", "application/json")
	w.Write(auth.RefreshToken(requestUser))
}
