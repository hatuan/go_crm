package controllers

import (
	"encoding/json"
	"erpvietnam/crm/auth"
	"erpvietnam/crm/models"
	"errors"
	ctx "github.com/gorilla/context"
	"net/http"
)

// ErrUsernameTaken is thrown when a user attempts to register a username that is taken.
var ErrRequestLoginInvalidate = errors.New("Request Login Invalidate")

func TokenAuth(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	switch {
	case r.Method == "POST": //login by token
		requestLogin := new(models.LoginDTO)
		err := json.NewDecoder(r.Body).Decode(&requestLogin)
		if err != nil {
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{ErrRequestLoginInvalidate.Error()}, IsAuthenticated: false}, http.StatusInternalServerError)
		}

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
