package service

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"os"
	"testing"

	"github.com/dannywolfmx/iwb/app/domain/entity"
)

var key *rsa.PrivateKey

func TestMain(m *testing.M) {
	var err error
	key, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Error on generate rsa key: %s", err)
	}
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestGenerateJWT(t *testing.T) {
	userTest := &entity.User{
		Name: "prueba",
	}

	service := NewTokenService(key, "https://localhost")

	_, err := service.GenerateJWT(userTest)
	if err != nil {
		t.Fatalf("Error on generate JWT: %s", err)
	}
}

func TestParseJWT(t *testing.T) {
	userTest := &entity.User{
		Name: "prueba",
	}

	service := NewTokenService(key, "https://localhost")

	payload, err := service.GenerateJWT(userTest)
	if err != nil {
		t.Fatalf("Error on generate JWT: %s", err)
	}

	_, err = service.ParseJWT(payload)

	if err != nil {
		t.Fatalf("Error on parse JWT: %s", err)
	}

}
