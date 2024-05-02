package generate

import (
	"crypto/rand"
	"math/big"
)

const CHARSET = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+"

// GeneratePassword generates a random password of a given length.
// The password is composed of characters from CHARSET.
// The function panics if there's an error generating random numbers.
func GeneratePassword(length int) string {
	// Create a byte slice to hold the password
	password := make([]byte, length)

	// For each index in the password
	for i := range password {
		// Set the character at the index to a random character from CHARSET
		password[i] = CHARSET[RandomInt(0, len(CHARSET)-1)]
	}

	// Convert the byte slice to a string and return it
	return string(password)
}

// RandomInt generates a random integer between min and max (inclusive) using crypto/rand.
// The function panics if max is less than min or if there's an error generating random numbers.
func RandomInt(min, max int) int {
	// Calculate the difference between max and min
	diff := max - min

	// If max is less than min, panic
	if diff < 0 {
		panic("max should be greater than min")
	}

	// Generate a random number between 0 and diff (inclusive)
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(diff+1)))
	if err != nil {
		// If there's an error generating the random number, panic
		panic(err)
	}

	// Convert the random number to an int64 and add min to it
	n := nBig.Int64() + int64(min)

	// Return the random number as an int
	return int(n)
}
