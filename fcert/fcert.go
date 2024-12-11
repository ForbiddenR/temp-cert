package fcert

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

func CalculateSignature(keyPath, filePath string) (string, error) {
	privateKey, err := loadPrivateKey(keyPath)
	if err != nil {
		return "", err
	}
	fileHash, err := hashFile(filePath)
	if err != nil {
		return "", err
	}
	fmt.Printf("hash: %x\n", fileHash)
	signature, err := signHash(fileHash, privateKey)
	if err != nil {
		return "", err
	}
	sig := base64.StdEncoding.EncodeToString(signature)
	return sig, nil
}

func loadPrivateKey(filePath string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "PRIVATE KEY" {
		fmt.Println(block.Type)
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	return privateKey.(*rsa.PrivateKey), err
}

func hashFile(filePath string) ([]byte, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	hash := sha256.New()
	_, err = hash.Write(fileData)
	if err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

func signHash(hash []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash)
}
