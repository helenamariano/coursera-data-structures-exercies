package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Exercise 2: Hashing with chains
// Implement a hash table with lists chaining.
func main() {
	numBuckets0 := 5
	queries0 := []string{
		"add world",
		"add HellO",
		"check 4",
		"find World",
		"find world",
		"del world",
		"check 4",
		"del HellO",
		"add luck",
		"add GooD",
		"check 2",
		"del good",
	}
	// Expected:
	// 	HellO world
	// 	no
	// yes
	// HellO
	// GooD luck
	fmt.Println("Query 0..")
	process(numBuckets0, queries0)

	numBuckets1 := 4
	queries1 := []string{
		"add test",
		"add test",
		"find test",
		"del test",
		"find test",
		"find Test",
		"add Test",
		"find Test",
	}
	// Expected:
	// yes
	// no
	// no
	// yes
	fmt.Println("Query 1..")
	process(numBuckets1, queries1)

	numBuckets2 := 3
	queries2 := []string{
		"check 0",
		"find help",
		"add help",
		"add del",
		"add add",
		"find add",
		"find del",
		"del del",
		"find del",
		"check 0",
		"check 1",
		"check 2",
	}
	// Expected:
	//
	// no
	// yes
	// yes
	// no
	//
	//add help
	//

	fmt.Println("Query 2..")
	process(numBuckets2, queries2)
}

func process(numBuckets int, queries []string) error {
	hashSet := NewHashSet(numBuckets)
	for _, query := range queries {
		q := strings.Split(query, " ")
		if len(q) != 2 {
			errMsg := fmt.Sprintf("Invalid query: %s\n", query)
			return errors.New(errMsg)
		}

		queryType := q[0]
		switch queryType {
		case "add":
			key := q[1]
			hashSet.Add(key)
		case "del":
			key := q[1]
			hashSet.Del(key)
		case "find":
			key := q[1]
			found := hashSet.Find(key)
			if found {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		case "check":
			bucket, _ := strconv.Atoi(q[1])
			hashSet.Check(bucket)
		default:
			errMsg := fmt.Sprintf("Unknown query type : %s\n", queryType)
			return errors.New(errMsg)
		}
	}
	return nil
}
