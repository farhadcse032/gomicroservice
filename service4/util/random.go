package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}
func RandomString(n int) string {
	var sb strings.Builder

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(len(alphabet))]
		sb.WriteByte(c)
	}
	return sb.String()
}
func RandomCurrency() string {
	currency := []string{"BDT", "USD", "EUR", "SGD", "INR"}
	return currency[rand.Intn(len(currency))]
}
