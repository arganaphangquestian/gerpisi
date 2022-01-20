package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserToken struct {
	AccessToken string `json:"access_token"`
}

func Verify(tokenString string) (*string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error : %s", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("token not valid")
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, fmt.Errorf("token can't be parsed")
	}
	return &claims.Id, nil
}

func CreateToken(userID string) (*UserToken, error) {
	accessToken := jwt.StandardClaims{}
	accessToken.Id = userID
	accessToken.ExpiresAt = time.Now().Add(time.Hour * 5).Unix() // 5 Hour
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessToken)
	aToken, err := at.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	return &UserToken{
		AccessToken: aToken,
	}, nil
}
