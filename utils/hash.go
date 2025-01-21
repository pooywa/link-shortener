package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func GenerateShortCode(url string) string {
	hash := sha256.Sum256([]byte(url))
	return base64.URLEncoding.EncodeToString(hash[:])[:8]
}
