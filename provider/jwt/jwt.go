package jwtprovider

import (
	"errors"
	"fmt"
	"iTask/common"
	"iTask/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// data to generate Token
type TokenPayload struct {
	Email string `json:"email"`
	Role  int    `json:"role"`
}

// data receive after generate
type Token struct {
	AccessToken string `json:"accessToken"`
	ExpiresAt   int64  `json:"expiresAt"`
}

type myClaim struct {
	Payload TokenPayload `json:"payload"`
	jwt.RegisteredClaims
}

func GenerateJWT(data TokenPayload, cfg *config.Config) (*Token, error) {
	expiredAt := jwt.NewNumericDate(time.Now().Local().Add(time.Hour * 12))
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim{
		data,
		jwt.RegisteredClaims{
			ExpiresAt: expiredAt,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
		},
	})

	//fmt.Println(cfg.App.Secret)

	accessToken, err := jwtToken.SignedString([]byte(cfg.App.Secret))

	if err != nil {
		return nil, err

	}
	return &Token{accessToken, expiredAt.Unix()}, nil
}

func ValidateJWT(accessToken string, cfg *config.Config) (*TokenPayload, error) {
	token, err := jwt.ParseWithClaims(accessToken, &myClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.App.Secret), nil
	})

	if err != nil {
		return nil, ErrInvalidToken
	}
	if !token.Valid {
		return nil, ErrInvalidToken
	}
	claims, ok := token.Claims.(*myClaim)

	if !ok {
		return nil, ErrInvalidToken
	}
	return &claims.Payload, nil
}

// Declare errors relate to token
var (
	ErrTokenNotFound = common.ErrAuthorized(errors.New("token not found"))
	ErrEncodingToken = common.ErrAuthorized(errors.New("error encoding token"))
	ErrInvalidToken  = common.ErrAuthorized(errors.New("invalid token"))
)
