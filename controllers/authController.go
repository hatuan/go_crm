package controllers

import (
	"encoding/json"
	"erpvietnam/crm/auth"
	"erpvietnam/crm/models"
	"net/http"
	ctx "github.com/gorilla/context"
)

func TokenAuth(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	switch {
	case r.Method == "POST": //login by token
		requestLogin := new(models.LoginDTO)
		json.NewDecoder(r.Body).Decode(&requestLogin)

		responseStatus, token := auth.TokenLogin(requestLogin)

		JSONResponse(w, token, responseStatus)
	case r.Method == "GET": //get user from token
		requestUser := ctx.Get(r, "user").(models.User)
		JSONResponse(w, requestUser, http.StatusOK)
	}

}

func TokenRefresh(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := ctx.Get(r, "user").(models.User)

	responseStatus, token := auth.TokenRefresh(requestUser.Name)

	JSONResponse(w, token, responseStatus)
}
