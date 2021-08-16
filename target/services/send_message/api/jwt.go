package api

import (
	"errors"
	"fmt"

	jwt "github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

func getJWTconfig(jwt_secret string) middleware.JWTConfig {
	signingKey := []byte(jwt_secret)

	config := middleware.JWTConfig{
		TokenLookup: "query:jwt_token",
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
				}
				return signingKey, nil
			}

			// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
			token, err := jwt.Parse(auth, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			return token, nil
		},
	}

	return config
}

func getJWTusername(jwt_secret string) string {
	return "TODO"
}
