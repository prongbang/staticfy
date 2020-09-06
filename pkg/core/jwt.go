package core

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

func Get(r *http.Request, key string) string {
	return GetStringBySecret(r, key, os.Getenv("JWT_SECRET"))
}

func GetStringBySecret(r *http.Request, key string, secretKey string) string {
	authorization := r.Header.Get("Authorization")
	if authorization != "" {
		tokenStr := strings.Replace(authorization, "Bearer ", "", -1)
		hmacSecret := []byte(secretKey)
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return hmacSecret, nil
		})

		if err != nil {
			return ""
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			uid := claims[key]
			return uid.(string)
		}
	}
	return ""
}
