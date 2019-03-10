package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

// Exercise 1: binary tree traversals
// You are given a rooted binary tree. Build and output its in-order, pre-order and post-order traversals.

// Exercise 2: Is it a binary search tree?

func main() {
	keys, left, right, err := parseArgs()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(-1)
	}

	bt := NewBinaryTree(keys, left, right)
	inOrder := bt.InOrder()
	fmt.Printf("%d\n", inOrder)
	preOrder := bt.PreOrder()
	fmt.Printf("%d\n", preOrder)
	postOrder := bt.PostOrder()
	fmt.Printf("%d\n", postOrder)
}

func parseArgs() (keys []int, left []int, right []int, err error) {

	args, err := getArgsFromCommandLine(-1)
	if err != nil {
		return nil, nil, nil, err
	}

	if len(args)%3 != 0 {
		return nil, nil, nil, errors.New("Invalid number of arguments")
	}

	ki := 0
	for i := 0; i < len(args); i += 3 {
		keys[ki] = args[i]
		left[ki] = args[i+1]
		right[ki] = args[i+2]
		ki++
	}

	return keys, left, right, nil
}

func getArgsFromCommandLine(expectedNumArgs int) (parsedArgs []int, err error) {
	flag.Parse()
	rawArgs := flag.Args()

	if expectedNumArgs > 1 && len(rawArgs) != expectedNumArgs {
		errMsg := fmt.Sprintf("Expected %d argument(s) got %d: \n%s", expectedNumArgs, len(rawArgs), rawArgs)
		err := errors.New(errMsg)
		return parsedArgs, err
	}

	for _, arg := range rawArgs {
		val, err := strconv.Atoi(arg)
		if err != nil {
			return parsedArgs, err
		}
		parsedArgs = append(parsedArgs, val)

	}
	return parsedArgs, nil
}
