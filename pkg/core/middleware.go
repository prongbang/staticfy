package core

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

func APIKeyMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-API-KEY")
		if key != "" && key == os.Getenv("X-API-KEY") {
			h.ServeHTTP(w, r)
		} else {
			Unauthorized(w)
		}
	}
}

func JWTMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		jwToken := strings.Replace(reqToken, "Bearer ", "", -1)
		token, err := jwt.Parse(jwToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("%s", "There was an error")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			Unauthorized(w)
			return
		}

		if token.Valid {
			h.ServeHTTP(w, r)
			return
		}

		Unauthorized(w)
	}
}
