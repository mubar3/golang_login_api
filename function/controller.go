package Controller

import (
	"crypto/sha256"
	"encoding/hex"

	_ "github.com/mattn/go-sqlite3"
)

type APIResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func isInSlice(slice []any, val any) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func hashSHA256(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashed := hasher.Sum(nil)
	return hex.EncodeToString(hashed)
}

func encryptHash(pass string, key string) string {
	encryptPass := hashSHA256(hashSHA256(pass) + key)
	return encryptPass
}
