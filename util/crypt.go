package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
)

func EncryptWithPubKey(msg []byte, pub *rsa.PublicKey) ([]byte, error) {
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}
