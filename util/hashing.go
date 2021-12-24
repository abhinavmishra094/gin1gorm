package util

import (
	"crypto/sha512"
	"encoding/base64"
	"os"
)

func HashPassword(password string) string {
	passwordBytes := []byte(password)
	sha512Hasher := sha512.New()
	salt := []byte(os.Getenv("SALT"))
	passwordBytes = append(passwordBytes, salt...)
	sha512Hasher.Write(passwordBytes)
	hashedPassword := sha512Hasher.Sum(nil)
	base64EncodedPasswordHash := base64.URLEncoding.EncodeToString(hashedPassword)
	return base64EncodedPasswordHash
}

func DoPasswordsMatch(hashedPassword, currPassword string) bool {

	currPasswordHash := HashPassword(currPassword)
	return hashedPassword == currPasswordHash

}
