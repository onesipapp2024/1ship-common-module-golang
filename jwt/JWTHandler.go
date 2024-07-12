package jwt

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"sync"
)

var privateKey *ed25519.PrivateKey
var publicKey *ed25519.PublicKey
var lockPrivateKey = &sync.Mutex{}
var lockPublicKey = &sync.Mutex{}

func GenerateKeyPair() (priKey string, pubKey string, err error) {
	pub, pri, err := ed25519.GenerateKey(nil)
	if err != nil {
		return "", "", err
	}
	priByte, err := x509.MarshalPKCS8PrivateKey(pri)
	if err != nil {
		return "", "", err
	}
	priBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: priByte,
	}
	priKey = string(pem.EncodeToMemory(priBlock))
	pubByte, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return "", "", err
	}
	pubBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubByte,
	}
	pubKey = string(pem.EncodeToMemory(pubBlock))
	return priKey, pubKey, nil
}

func GetPrivateKey() (*ed25519.PrivateKey, error) {
	if privateKey == nil {
		lockPrivateKey.Lock()
		defer lockPrivateKey.Unlock()
		if privateKey == nil {
			priKey := viper.GetString("JWT_PRIVATE_KEY")
			if priKey == "" {
				err := errors.New("JWT_PRIVATE_KEY is not set")
				return nil, err
			}
			priBlock, _ := pem.Decode([]byte(priKey))
			if priBlock == nil {
				err := errors.New("can not decode JWT_PRIVATE_KEY")
				return nil, err
			}
			key, err := x509.ParsePKCS8PrivateKey(priBlock.Bytes)
			if err != nil {
				return nil, err
			}
			convertedKey := key.(ed25519.PrivateKey)
			privateKey = &convertedKey
		}
	}
	return privateKey, nil
}

func GetPublicKey() (*ed25519.PublicKey, error) {
	if publicKey == nil {
		lockPublicKey.Lock()
		defer lockPublicKey.Unlock()
		if publicKey == nil {
			pubKey := viper.GetString("JWT_PUBLIC_KEY")
			if pubKey == "" {
				err := errors.New("JWT_PUBLIC_KEY is not set")
				return nil, err
			}
			pubBlock, _ := pem.Decode([]byte(pubKey))
			if pubBlock == nil {
				err := errors.New("can not decode JWT_PUBLIC_KEY")
				return nil, err
			}
			key, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
			if err != nil {
				return nil, err
			}
			convertedKey := key.(ed25519.PublicKey)
			publicKey = &convertedKey
		}
	}
	return publicKey, nil
}

func ParseJWT(token string, claims jwt.Claims) (*jwt.Token, error) {
	key, err := GetPublicKey()
	if err != nil {
		return nil, err
	}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(jwtToken *jwt.Token) (interface{}, error) {
		return *key, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodEdDSA.Alg()}))
	return jwtToken, err
}
