package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_tokenService "go-todo-list/authentication/service"
	"go-todo-list/lib"
	"net/http"
	"strings"
)

func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			lib.CreateResponse(401, 401, "Token is not valid", nil).JSON(w)
			return
		}

		secret := viper.GetString("application.jwt.secret")
		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		claims := &_tokenService.JWtClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return []byte(secret), nil
		})

		if err != nil {
			logrus.Error(err)
			lib.CreateResponse(401, 401, fmt.Sprintf("Error while parsing Token: %s", tokenString), nil).JSON(w)
			return
		}

		if !token.Valid {
			lib.CreateResponse(401, 401, fmt.Sprintf("Token is not valid: %s", tokenString), nil).JSON(w)
			return
		}

		context.Set(r, "userId", token)
		next.ServeHTTP(w, r)
	})
}

