package token

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct{}

func NewToken() *Token {
	return &Token{}
}

func (t *Token) GenerateToken(id int) (string, error) {
	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func (t *Token) extractTokenString(bearToken string) string {
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

func (t *Token) parseTokenAndCheckExpired(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (t *Token) ExtractUserId(bearToken string) (int, error) {
	tokenString := t.extractTokenString(bearToken)

	token, err := t.parseTokenAndCheckExpired(tokenString)
	if err != nil {
		return -1, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, ok := claims["user_id"].(float64)
		if !ok {
			return -1, errors.New("Invalid user_id in specified token")
		}

		return int(userId), nil
	}

	return -1, errors.New("Specified token is invalid")
}
