package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Exercies 1: Phone book
// Implement a simple phone book manager.

func main() {

	queries0 := []string{
		"add 911 police",
		"add 76213 Mom",
		"add 17239 Bob",
		"find 76213",
		"find 910",
		"find 911",
		"del 910",
		"del 911",
		"find 911",
		"find 76213",
		"add 76213 daddy",
		"find 76213",
	}

	fmt.Println("Processing query 0...")
	if process(queries0) != nil {
		os.Exit(-1)
	}

	queries1 := []string{
		"find 3839442",
		"add 123456 me",
		"add 0 granny",
		"find 0",
		"find 123456",
		"del 0",
		"del 0",
		"find 0",
	}

	fmt.Println("Processing query 1...")
	if process(queries1) != nil {
		os.Exit(-1)
	}

}

func process(queries []string) error {

	errMsg := ""
	phoneBook := NewPhoneBook()
	for _, query := range queries {
		q := strings.Split(query, " ")
		queryType := q[0]

		switch queryType {
		case "add":
			if len(q) != 3 {
				errMsg = fmt.Sprintf("Bad query type: %s", q)
				break
			}
			number, _ := strconv.Atoi(q[1])
			name := q[2]
			phoneBook.Add(number, name)
		case "find":
			if len(q) != 2 {
				errMsg = fmt.Sprintf("Bad query type: %s", q)
				break
			}
			number, _ := strconv.Atoi(q[1])
			name := phoneBook.Find(number)
			fmt.Println(name)
		case "del":
			if len(q) != 2 {
				errMsg = fmt.Sprintf("Bad query type: %s", q)
				break
			}
			number, _ := strconv.Atoi(q[1])
			phoneBook.Del(number)
		default:
			errMsg = fmt.Sprintf("Unknown query type: %s", queryType)
			break
		}
	}

	if errMsg != "" {
		return errors.New(errMsg)
	}

	return nil
}
