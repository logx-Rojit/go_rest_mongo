package utils

import (
	"math/rand"
	"time"
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

func GenerateRandomPassword(size int8, useLetters bool, useSpecialChar bool, useNumberChar bool) string {
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
	return string(b)

}
