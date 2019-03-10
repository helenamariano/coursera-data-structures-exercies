package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numWorkers, jobs, err := parseArgs()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(-1)
	}

	result := Solve(numWorkers, jobs)
	for _, r := range result {
		fmt.Printf("%d %d\n", r[0], r[1])
	}

}

func parseArgs() (int, []int, error) {

	args, err := getArgsFromCommandLine(-1)
	if err != nil {
		return 0, nil, err
	}

	if len(args) < 2 {
		return 0, nil, errors.New("Invalid number of arguments")
	}

	numWokers := args[0]
	numJobs := args[1]
	if len(args) != numJobs+2 {
		return 0, nil, errors.New("Invalid number of jobs")
	}

	jobs := args[2:]
	return numWokers, jobs, nil
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
