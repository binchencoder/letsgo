package strings

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var letterNumRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var numCapitalRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var hexRunes = []rune("0123456789ABCDEF")
var numberRunes = []rune("0123456789")

func randRunes(n int, runes []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

// RandLetters generates random lower and upper case letters with length specified.
func RandLetters(n int) string {
	return randRunes(n, letterRunes)
}

// RandString generates random lower, upper case letters and numbers with length
// specified.
func RandString(n int) string {
	return randRunes(n, letterNumRunes)
}

// RandHex generates random numbers in Hex format with length specified.
func RandHex(n int) string {
	return randRunes(n, hexRunes)
}

// RandNumCapitals generates random upper case letters and numbers with length
// specified.
func RandNumCapitals(n int) string {
	return randRunes(n, numCapitalRunes)
}

// RandNumbers generates random numbers with length specified.
func RandNumbers(n int) string {
	return randRunes(n, numberRunes)
}
