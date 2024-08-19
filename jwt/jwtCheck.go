package jwt

import (
	"errors"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func AuthToken(username string, token string) (bool, error) {
	SECRET_KEY := os.Getenv("SECRET_KEY")

	if SECRET_KEY == "" {
		err := errors.New("SECRET KEY not found in env")
		return false, err
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return false, err
	}

	if !parsedToken.Valid {
		return false, nil
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return false, fmt.Errorf("failed to parse claims")
	}

	claimedUsername, ok := claims["username"].(string)
	if !ok {
		return false, fmt.Errorf("invalid username claim")
	}

	if claimedUsername != username {
		return false, nil
	}

	return true, nil
}
