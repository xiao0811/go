package handlers

import (
	"crypto/sha256"
	"encoding/hex"
)

func MYSHA256(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	hashValue := hash.Sum(nil)
	return hex.EncodeToString(hashValue)
}
