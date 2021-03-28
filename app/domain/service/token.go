package service

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"

	"github.com/dannywolfmx/iwb/app/domain/entity"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

type TokenService struct{}

func (t *TokenService) GenerateJWT(user *entity.User) (entity.Token, error) {
	issuer := "github.com/dannywolfmx/iwb/app/domain/entity"

	token := jwt.New()

	token.Set(jwt.IssuerKey, issuer)

	//Encript
	//
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", err
	}

	signed, err := jwt.Sign(token, jwa.RS256, key)

	if err != nil {
		return "", err
	}
	//Encode to base64
	return base64.StdEncoding.EncodeToString(signed), nil
}
