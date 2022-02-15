package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(stringSize int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < stringSize; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomCurrency() string {
	currencies := []string{"EUR", "INR", "USD", "AUD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
