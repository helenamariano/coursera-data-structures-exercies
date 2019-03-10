package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// Exercise 1
// Check brackets in the code
// Input contains one string 𝑆 which consists of big and small latin letters, digits, punctuation marks and brackets from the set []{}().
// Output: S uses brackets correctly, output true, followed by index at where 1st matching bracket failed

type bracketItem struct {
	value rune
	index int
}

func main() {
	input, err := parseArguments()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}

	fmt.Printf("Input: %s\n", input)
	isBal, index, err := isBalanced(input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}
	fmt.Printf("Result: %v\n", isBal)
	fmt.Printf("Index: %v\n", index)

}

func isBalanced(s string) (bool, int, error) {
	if len(s) == 0 {
		return false, 0, errors.New("String is empty")
	}
	stack := StackNew()

	for i, c := range s {
		if isOpenBracket(c) {
			openBracket := bracketItem{value: c, index: i + 1}
			stack.Push(openBracket)
		}

		if isClosedBracket(c) {
			if stack.Empty() {
				return false, i + 1, nil
			}

			openBracket, _ := stack.Pop()
			if !doesBracketMatch(openBracket.(bracketItem).value, c) {
				return false, i + 1, nil
			}
		}
	}

	if stack.Empty() {
		return true, 0, nil
	}

	openBracket, _ := stack.Top()
	return false, openBracket.(bracketItem).index, nil
}

func isOpenBracket(c rune) bool {
	return (c == '[') || (c == '(') || (c == '{')
}

func isClosedBracket(c rune) bool {
	return (c == ']') || (c == ')') || (c == '}')
}

func doesBracketMatch(openBracket, closedBracket rune) bool {
	return (openBracket == '(' && closedBracket == ')') || (openBracket == '[' && closedBracket == ']') || (openBracket == '{' && closedBracket == '}')
}

func parseArguments() (string, error) {
	flag.Parse()
	if len(flag.Args()) != 1 {
		return "", errors.New("Invalid number of arguments")
	}

	return os.Args[0], nil
}
