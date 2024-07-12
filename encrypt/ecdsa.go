package encrypt

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func ParseECDSAPrivateKey(privateKey *ecdsa.PrivateKey, priKey string) error {
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
	convertedKey := key.(*ecdsa.PrivateKey)
	*privateKey = *convertedKey
	return nil
}

func ParseECDSAPublicKey(publicKey *ecdsa.PublicKey, pubKey string) error {
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
	convertedKey := key.(*ecdsa.PublicKey)
	*publicKey = *convertedKey
	return nil
}

func SignWithECDSA(privateKey *ecdsa.PrivateKey, message string) (string, error) {
	hash := sha256.Sum256([]byte(message))
	sign, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", err
	}
	return string(sign), nil
}

func VerifyECDSASign(publicKey *ecdsa.PublicKey, message string, sign string) bool {
	hash := sha256.Sum256([]byte(message))
	return ecdsa.VerifyASN1(publicKey, hash[:], []byte(sign))
}
