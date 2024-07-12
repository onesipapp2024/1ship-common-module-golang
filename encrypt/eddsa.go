package encrypt

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func ParseEDDSAPrivateKey(privateKey *ed25519.PrivateKey, priKey string) error {
	if priKey == "" {
		err := errors.New("private key is not valid")
		return err
	}
	priBlock, _ := pem.Decode([]byte(priKey))
	if priBlock == nil {
		err := errors.New("can not decode private key")
		return err
	}
	key, err := x509.ParsePKCS8PrivateKey(priBlock.Bytes)
	if err != nil {
		return err
	}
	convertedKey := key.(ed25519.PrivateKey)
	privateKey = &convertedKey
	return nil
}

func ParseEDDSAPublicKey(publicKey *ed25519.PublicKey, pubKey string) error {
	if pubKey == "" {
		err := errors.New("public key is not valid")
		return err
	}
	pubBlock, _ := pem.Decode([]byte(pubKey))
	if pubBlock == nil {
		err := errors.New("can not decode public key")
		return err
	}
	key, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return err
	}
	convertedKey := key.(ed25519.PublicKey)
	publicKey = &convertedKey
	return nil
}

func SignWithEDDSA(privateKey *ed25519.PrivateKey, message string) string {
	return string(ed25519.Sign(*privateKey, []byte(message)))
}

func VerifyEDDSASign(publicKey *ed25519.PublicKey, message string, sign string) bool {
	return ed25519.Verify(*publicKey, []byte(message), []byte(sign))
}
