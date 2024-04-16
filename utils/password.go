package utils

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	initailString = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	specialChar   = "!@#$%^&*()_+{}:',./<>?"
	numberChar    = "1234567890"
)

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)
}

func GenerateRandomPassword(size int8, useLetters bool, useSpecialChar bool, useNumberChar bool) (string, string) {
	b := make([]byte, size)

	for i := range b {
		if useLetters {
			b[i] = initailString[rand.Intn(len(initailString))]
		} else if useSpecialChar {
			b[i] = specialChar[rand.Intn(len(specialChar))]
		} else if useNumberChar {
			b[i] = numberChar[rand.Intn(len(numberChar))]
		}
	}
	hash, _ := HashPassword(string(b))
	return hash, string(b)
}

func HashPassword(password string) (string, error) {
	salt := 14
	hash, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(hash), err
}

func ComparePassword(password string, dbPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}
