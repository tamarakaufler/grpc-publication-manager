package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	proto "github.com/tamarakaufler/grpc-publication-manager/author-service/proto"
)

var (
	// Secure key string used as a salt when hashing tokens
	key = []byte("b9d4d807c28ca2b63bc7468e350e4493")
)

// hashed and sent as the second segment of the JWT
type CustomClaims struct {
	Author *proto.Author
	jwt.StandardClaims
}

type Authentication interface {
	Decode(token string) (*CustomClaims, error)
	Encode(author *proto.Author) (string, error)
}

type TokenService struct {
}

func (ts TokenService) Decode(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	// Validate the token => return custom claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func (ts TokenService) Encode(author *proto.Author) (string, error) {
	expireToken := time.Now().Add(time.Hour * 48).Unix()

	claims := CustomClaims{
		author,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "publication.manager.author",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
