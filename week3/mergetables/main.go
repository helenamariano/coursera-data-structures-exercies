package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

// Exercise 3: Merge tables
// Simulate a sequence of merge operations with tables in a database
// Input: n (num tables in database), m (num merges), r-1...r-n (number of rows in ith table)
// dest-1, source-1... dest-m, source-m

func main() {
	tables, merges, err := parseArgs()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(-1)
	}
	fmt.Printf("Tables: %v\n", tables)
	fmt.Printf("Merges: %v\n", merges)

	disjointSet := NewDisjointSet(tables)
	for i := 0; i < len(merges); i += 2 {
		destination := merges[i]
		source := merges[i+1]
		// inputs are indexed from one but are stored from index 0
		maxTableSize := disjointSet.Merge(destination-1, source-1)
		fmt.Printf("(%d, %d) => %d\n", destination, source, maxTableSize)
	}
}

func parseArgs() ([]int, []int, error) {

	args, err := getArgsFromCommandLine(-1)
	if err != nil {
		return nil, nil, err
	}

	if len(args) < 2 {
		return nil, nil, errors.New("Invalid number of arguments")
	}

	numTables := args[0]
	numMerges := args[1]
	if len(args) != numTables+(numMerges*2)+2 {
		return nil, nil, errors.New("Invalid number of arguments")
	}

	tables := args[2 : numTables+2]
	merges := args[numTables+2:]
	return tables, merges, nil
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
