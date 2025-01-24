package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var ownerNames = [50]string{
	"Alex", "Emma", "Liam", "Sophia", "Noah", "Olivia", "Ethan", "Ava", "Mason", "Isabella",
	"Lucas", "Mia", "Logan", "Amelia", "James", "Harper", "Benjamin", "Evelyn", "Henry", "Charlotte",
	"William", "Abigail", "Oliver", "Ella", "Elijah", "Scarlett", "Michael", "Emily", "Daniel", "Elizabeth",
	"Jackson", "Luna", "Sebastian", "Chloe", "Carter", "Grace", "Alexander", "Victoria", "Mateo", "Zoey",
	"Matthew", "Penelope", "David", "Riley", "Joseph", "Layla", "Samuel", "Lily", "Leo", "Aria",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return ownerNames[RandomInt(0, 49)]
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generate a random currency code
func RandomCurrency() string {
	currencies := []string{"EUR", "USD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
