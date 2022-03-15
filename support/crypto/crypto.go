package crypto

import (
	"api/support"
	"crypto/md5"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func Hash(pw string) string {
	hash := md5.Sum([]byte(pw))
	return hex.EncodeToString(hash[:])
}

func Hash2(pw string) string {
	password := []byte(pw)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	support.Panic(err)

	return string(hashedPassword)
}

func isValid(pw string, hashed string) bool {
	psByte := []byte(pw)
	hsByte := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(hsByte, psByte)
	if err == nil {
		return true
	}
	return false
}
