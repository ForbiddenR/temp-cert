package fcert

import (
	"crypto/rsa"
	"os"
)

func loadPrivateKey(filePath string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(filePath)
}