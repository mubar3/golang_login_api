package Package

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
)

func hashSHA256(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashed := hasher.Sum(nil)
	return hex.EncodeToString(hashed)
}

func EncryptHash(pass string, key string) string {
	encryptPass := hashSHA256(hashSHA256(pass) + key)
	return encryptPass
}

// Fungsi Validasi untuk memeriksa data form
func Validasi(data map[string][]string, rules map[string]string) (bool, error) {
	for field, rule := range rules {
		if rule == "required" {
			if values, ok := data[field]; !ok || len(values[0]) == 0 {
				return false, errors.New("Inputan " + field + " kosong")
			}
		} else if rule == "numeric" {
			values, ok := data[field]
			if !ok || len(values[0]) == 0 {
				return false, errors.New("Inputan " + field + " kosong")
			}

			// Coba konversi nilai string ke integer
			_, err := strconv.Atoi(values[0])
			if err != nil {
				return false, errors.New("Inputan " + field + " bukan angka")
			}
		}
	}
	return true, nil
}
