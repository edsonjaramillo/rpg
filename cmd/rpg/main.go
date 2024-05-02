package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"rpg/internal/generate"

	"github.com/atotto/clipboard"
)

func main() {
	length := GetNumberFromArgs()
	password := generate.GeneratePassword(length)
	err := clipboard.WriteAll(password)
	if err != nil {
		fmt.Println("Failed to copy the password to the clipboard.")
		panic(err)
	}
	fmt.Println(password)
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

// Snv8Q59q1*e3w*VS KjuHF66ruYualybr
