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
		TokenLookup: "query:jwt_key",
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("Unexpected jwt signing method=%v", t.Header["alg"])
				}
				return signingKey, nil
			}

			// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
			token, err := jwt.Parse(auth, keyFunc)

			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("Invalid token")
			}
			_, ok := token.Claims.(jwt.Claims)
			if !ok {
				return nil, errors.New("No claims")
			}

			return token, nil
		},
	}

	return config
}

func extractUsername(jwt_token string, jwt_secret string) string {
	signingKey := []byte(jwt_secret)

	keyFunc := func(t *jwt.Token) (interface{}, error) {
		return signingKey, nil
	}

	token, _ := jwt.Parse(jwt_token, keyFunc)
	claims, _ := token.Claims.(jwt.MapClaims)

	return claims["username"].(string)
}

func hasChannel(channel string, jwt_token string, jwt_secret string) bool {
	signingKey := []byte(jwt_secret)

	keyFunc := func(t *jwt.Token) (interface{}, error) {
		return signingKey, nil
	}

	token, _ := jwt.Parse(jwt_token, keyFunc)
	claims, _ := token.Claims.(jwt.MapClaims)

	channels := claims["channels"].([]interface{})
	for _, elmt := range channels {
		if channel == elmt {
			return true
		}
	}
	return false
	/*
		return claims["channels"].([]string())
		return claims["channels"].([]interface{})
		return claims["channels"].([]string)
		return []string{"840656308797833227"}
	*/
}
