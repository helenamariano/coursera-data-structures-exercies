package main

import (
	"coursera-data-structures-exercies/week1/treeheight/tree"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {

	args, err := getArgsFromCommandLine(-1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}

	numNodes := args[0]
	parents := args[1:]

	var nodes []*tree.Node

	for i := 0; i < numNodes; i++ {
		node := tree.CreateNode(i)
		nodes = append(nodes, node)
	}

	fmt.Printf("Child/Parent link: %v", parents)

	for child, parent := range parents {
		// not a root node
		if parent >= 0 {
			nodes[child].SetParent(nodes[parent])
		}
	}

	tree.PrintTree(nodes)
	height := nodes[0].HeightUsingRecursion()
	fmt.Printf("Height: %d\n", height)
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
