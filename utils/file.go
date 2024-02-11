package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

func generateKey() (*ecdsa.PrivateKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	err = savePrivateKeyToFile(privateKey, fileName)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func savePrivateKeyToFile(privateKey *ecdsa.PrivateKey, filename string) error {
	keyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return err
	}

	keyFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer keyFile.Close()

	if err := pem.Encode(keyFile, &pem.Block{Type: "EC PRIVATE KEY", Bytes: keyBytes}); err != nil {
		return err
	}
	return nil
}

func readPrivateKeyFromFile(filename string) (*ecdsa.PrivateKey, error) {
	keyFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyFile)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func extractECDSAPublicKey(privateKey *ecdsa.PrivateKey) *ecdsa.PublicKey {
	return &privateKey.PublicKey
}
