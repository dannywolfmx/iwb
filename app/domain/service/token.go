package service

import (
	"crypto/rsa"
	"encoding/base64"
	"errors"

	"github.com/dannywolfmx/iwb/app/domain/entity"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

const (
	userkey = "user"
)

type TokenService struct {
	key    *rsa.PrivateKey
	issuer string
}

//NewTokenService generate a new token jwt
func NewTokenService(key *rsa.PrivateKey, issuer string) *TokenService {
	return &TokenService{
		key:    key,
		issuer: issuer,
	}
}

func (t *TokenService) GenerateJWT(user *entity.User) (entity.Token, error) {
	token := jwt.New()

	//Set key value to the jwt
	token.Set(jwt.IssuerKey, t.issuer)
	token.Set(userkey, user.Name)
	//Encript
	//

	signed, err := jwt.Sign(token, jwa.RS256, t.key)

	if err != nil {
		return "", err
	}
	//Encode to base64
	return base64.StdEncoding.EncodeToString(signed), nil
}

func (t *TokenService) ParseJWT(encodePayload string) (*entity.User, error) {
	payload, err := base64.StdEncoding.DecodeString(encodePayload)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(payload, jwt.WithValidate(true), jwt.WithVerify(jwa.RS256, &t.key.PublicKey))

	if err != nil {
		return nil, err
	}
	userRaw, ok := token.Get(userkey)
	if !ok {
		return nil, errors.New("User can't be finded")
	}
	user, ok := userRaw.(string)

	if !ok {
		return nil, errors.New("Can't cast to string the usecase name")
	}

	return &entity.User{
		Name: user,
	}, nil
}
