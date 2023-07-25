package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func EncryptWithPubKey(msg []byte, pub *rsa.PublicKey) ([]byte, error) {
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}

func ParsePubKey(key string) (*rsa.PublicKey, error) {
	pubPem, _ := pem.Decode([]byte(key))
	if pubPem == nil {
		return nil, errors.New("could not decode key")
	}
	if pubPem.Type != "RSA PUBLIC KEY" {
		return nil, errors.New("invalid pub key")
	}
	pubPemBytes := pubPem.Bytes
	parsedPubKey, err := x509.ParsePKCS1PublicKey(pubPemBytes)
	if err != nil {
		return nil, err
	}

	return parsedPubKey, nil
}

func PubKeyToString(key *rsa.PublicKey) string {
	pubkeyBytes, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return ""
	}
	pubkeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkeyBytes,
		},
	)

	return string(pubkeyPem)
}
