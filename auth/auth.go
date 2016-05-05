package auth

import (
	"encoding/json"
	"erpvietnam/crm/models"
	"net/http"
)

// Login attempts to login the user given a request.
func Login(requestLogin *models.LoginDTO) (int, []byte) {
	authBackend := InitJWTAuthenticationBackend()

	if authBackend.Authenticate(requestLogin.UserName, requestLogin.Password) {
		token, err := authBackend.GenerateToken(requestLogin.UserName)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		}
		response, _ := json.Marshal(models.Token{token})
		return http.StatusOK, response
	}

	return http.StatusUnauthorized, []byte("")
}

//RefreshToken get new JWT Token
func RefreshToken(requestUser *models.LoginDTO) []byte {
	authBackend := InitJWTAuthenticationBackend()
	token, err := authBackend.GenerateToken(requestUser.UserName)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(models.Token{token})
	if err != nil {
		panic(err)
	}
	return response
}
