package str

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const letterAndNumBytes = letterBytes + "0123456789"

// RandomGenerateChars -> it will generate based on chars only! (alphachars)
func RandomGenerateChars(length int) string {
	// Generate the seed... if we are not generating it, it will be the same value all the time!
	rand.Seed(time.Now().UnixNano())
	// Generate the random value
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// RandomGenerate -> it will generate based on chars & numbers only! (alphanumeric)
func RandomGenerate(length int) string {
	// Generate the seed... if we are not generating it, it will be the same value all the time!
	rand.Seed(time.Now().UnixNano())
	// Generate the random value
	b := make([]byte, length)
	for i := range b {
		b[i] = letterAndNumBytes[rand.Intn(len(letterAndNumBytes))]
	}
	return string(b)
}
