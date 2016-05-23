package auth

import (
	"erpvietnam/crm/log"
	"erpvietnam/crm/models"
	"net/http"
	"errors"
)

// ErrLoginInvalid is thrown when a user attempts to register a username that is taken.
var ErrLoginInvalid = errors.New("Login invalid.")

// TokenLogin attempts to login the user given a request.
func TokenLogin(requestLogin *models.LoginDTO) (int, models.Token) {
	authBackend := InitJWTAuthenticationBackend()

	if authBackend.Authenticate(requestLogin.UserName, requestLogin.Password) {
		token, err := authBackend.GenerateToken(requestLogin.UserName)
		if err != nil {
			log.Error(err)
			token := models.Token{
				TransactionalInformationDTO: models.TransactionalInformationDTO{ReturnMessage: [] string{err.Error()}, ReturnStatus: false},
				Token : "",
			}
			return http.StatusInternalServerError, token
		}
		response := models.Token{
			TransactionalInformationDTO: models.TransactionalInformationDTO{ReturnStatus: true},
			Token: token,
		}
		return http.StatusOK, response
	}

	return http.StatusUnauthorized, models.Token{
		TransactionalInformationDTO: models.TransactionalInformationDTO{ReturnMessage:  [] string{ErrLoginInvalid.Error()}, ReturnStatus: false},
		Token: "",
	}
}

//TokenRefresh get new JWT Token
func TokenRefresh(username string) (int, models.Token) {
	authBackend := InitJWTAuthenticationBackend()
	token, err := authBackend.GenerateToken(username)
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, models.Token{
			TransactionalInformationDTO: models.TransactionalInformationDTO{ReturnMessage: [] string{err.Error()}, ReturnStatus: false},
			Token: "",
		}
	}
	response := models.Token{
		TransactionalInformationDTO: models.TransactionalInformationDTO{ReturnStatus: true},
		Token: token,
	}
	return http.StatusOK, response
}
