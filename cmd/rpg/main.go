package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"

	"github.com/atotto/clipboard"
)

func main() {
	length := GetNumberFromArgs()
	password := GeneratePassword(length)
	err := clipboard.WriteAll(password)
	if err != nil {
		fmt.Println("Failed to copy the password to the clipboard.")
		panic(err)
	}
	fmt.Println("Password copied to clipboard.")
}

func GetNumberFromArgs() int {
	// Regex to match positive integers only
	numberRegex := regexp.MustCompile(`^[1-9]\d*$`)

	// Check if a command-line argument was provided
	if len(os.Args) > 1 {
		arg := os.Args[1]
		// If the argument matches the regex, convert it to an integer and return it
		if numberRegex.MatchString(arg) {
			number, _ := strconv.Atoi(arg)
			return number
		}
	}

	// Default to 12 if no valid argument was provided
	return 12
}

// GeneratePassword generates a random password of the specified length.
// The password can contain uppercase and lowercase letters, digits, and special characters.
func GeneratePassword(length int) string {
	// Define the characters that can be in the password
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789" +
		"!@#$%^&*()_+"

	// Create a byte slice of the specified length
	password := make([]byte, length)

	// For each byte in the slice, assign a random character from the charset
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	// Convert the byte slice to a string and return it
	return string(password)
}
