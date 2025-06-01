package middleware

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

var TokenAuth = jwtauth.New("HS256", []byte("key"), nil)

func Verifier() func(http.Handler) http.Handler{
	return jwtauth.Verifier(TokenAuth)
}

func Authenticator() func (http.Handler) http.Handler {
	return jwtauth.Authenticator(TokenAuth)
}

func GetUserIDFromContext(r *http.Request)(int, bool){
	_, claims, _ := jwtauth.FromContext(r.Context())
	idFloat, ok := claims["user_id"].(float64)
	return int(idFloat), ok
}