package util

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt returns a random integer between min and max.
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString returns a random string of length n.
func RandomString(n int) string {

	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}

// RandomEmail returns a random email address.
func RandomEmail() string {
	return RandomString(6) + "@test.com"
}

// RandomName returns a random name.
func RandomName() string {
	return RandomString(6)
}

// RandomPassword returns a random password of length n using numbers and letters.
func RandomPassword(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(b)
}

// Random Money returns a random amount of money.
func RandomMoney() int64 {
	return RandomInt(1, 1000)
}
