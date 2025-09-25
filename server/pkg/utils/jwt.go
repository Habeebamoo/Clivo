package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/Habeebamoo/Clivo/server/internal/config"
	"github.com/Habeebamoo/Clivo/server/internal/models"

	"github.com/golang-jwt/jwt/v5"
)


func SignToken(payload models.TokenPayload) (string, error) {
	JWT_KEY, err := config.Get("JWT_KEY")
	if err != nil {
		log.Fatal(err)
	}

	claims := jwt.MapClaims{
		"user_id": payload.UserId,
		"role": payload.Role,
		"exp": time.Now().Add(1*time.Hour).Unix(),
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenString.SignedString([]byte(JWT_KEY))
}

func ParseToken(token string) (models.TokenPayload ,error) {
	JWT_KEY, err := config.Get("JWT_KEY")
	if err != nil {
		log.Fatal(err)
	}

	//parse jwt
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", jwt.ErrSignatureInvalid
		}

		return []byte(JWT_KEY), nil
	})

	//error check
	if err != nil {
		return models.TokenPayload{}, fmt.Errorf("failed to verify token")
	}

	if !parsedToken.Valid {
		return models.TokenPayload{}, fmt.Errorf("invalid token")
	}

	//extract claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return models.TokenPayload{}, fmt.Errorf("invalid token payload")
	}

	//validate expiration
	exp, ok := claims["exp"].(float64)
	if !ok {
		return models.TokenPayload{}, fmt.Errorf("invalid token payload")
	}

	if time.Unix(int64(exp), 0).Before(time.Now()) {
		return models.TokenPayload{}, fmt.Errorf("token expired")
	}

	//extract user details from claims
	userId := claims["userId"].(string)
	role := claims["role"].(string)

	userDetails := models.TokenPayload{
		UserId: userId,
		Role: role,
	}

	return userDetails, nil
}