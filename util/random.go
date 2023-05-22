package util

import (
	"math/rand"
	"strings"
	"time"
)
const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixMicro()))
}

// Generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // return random in between 0 -> max-min then add min to expression 
}

// Generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// Generates a random owner name
func RandomOwnerName() string {
  return RandomString(10)
}

// Generates a random amount of money
func RandomMoney() int64 {
  return RandomInt(0, 1000)
}

// Generates a random currency
func RandomCurrency() string {
	currencyOptions := []string{"USD", "EUR", "GBP", "MXN", "JPY"}
	n := len(currencyOptions)
	return currencyOptions[rand.Intn(n)]
}