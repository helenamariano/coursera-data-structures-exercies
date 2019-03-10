package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// Exercise 3: Find pattern in text
// Implement the Rabin–Karp’s algorithm for searching the given pattern in the given text.
// Input: pattern, text
// Output: all the positions of the occurrences of pattern in text in the ascending order

func main() {

	pattern, text, err := getArgsFromCommandLine(2)
	if err != nil {
		fmt.Printf("Error: \n")
		os.Exit(-1)
	}

	fmt.Printf("Pattern: %s\n", pattern)
	fmt.Printf("Text: %s\n", text)
	result := FindPattern(pattern, text)
	fmt.Printf("Result: %v\n", result)

}

func getArgsFromCommandLine(expectedNumArgs int) (pattern, text string, err error) {
	flag.Parse()
	rawArgs := flag.Args()

	if expectedNumArgs > 1 && len(rawArgs) != expectedNumArgs {
		errMsg := fmt.Sprintf("Expected %d argument(s) got %d: \n%s", expectedNumArgs, len(rawArgs), rawArgs)
		err = errors.New(errMsg)
		return pattern, text, err
	}

	pattern = rawArgs[0]
	text = rawArgs[1]

	return pattern, text, nil
}
