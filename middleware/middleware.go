package middleware

import (
	"erpvietnam/crm/auth"
	"fmt"
	"net/http"

	"encoding/json"
	"erpvietnam/crm/models"
	jwt "github.com/dgrijalva/jwt-go"
	ctx "github.com/gorilla/context"
)

type Context struct {
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	next(rw, req)
	// Remove context contents
	ctx.Clear(req)
}

func RequireTokenAuthentication(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	authBackend := auth.InitJWTAuthenticationBackend()

	token, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return authBackend.PublicKey, nil
	})

	if err == nil && token.Valid {
		user := models.User{}
		user_bytes := []byte(token.Claims["user"].(string))
		json.Unmarshal(user_bytes, &user)
		ctx.Set(req, "user", user)

		next(rw, req)
	} else {
		transactionalInformation := new(models.TransactionalInformation)
		transactionalInformation.IsAuthenticated = false
		transactionalInformation.ReturnStatus = false
		transactionalInformation.ReturnMessage = []string{"Auth failed"}

		response, err := json.Marshal(transactionalInformation)
		if err != nil {
			response = []byte("{\"return_message\":\"" + err.Error() + "\"}\"")
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write(response)
	}
}
