package main

import "math/rand"

const prime = int64(1000000007)

func FindPattern(pattern, text string) []int {
	result := []int{}
	x := int64(0)
	for x == 0 {
		x = rand.Int63n(prime)
	}

	patternHash := polyHash(pattern, x)
	for i := 0; i < len(text)-len(pattern)+1; i++ {
		s := text[i : i+len(pattern)]
		sHash := polyHash(s, x)
		if patternHash != sHash {
			continue
		}
		if pattern == s {
			result = append(result, i)
		}
	}
	return result
}

func polyHash(s string, x int64) int64 {
	result := int64(0)
	for i := len(s) - 1; i >= 0; i-- {
		result = ((result * x) + int64(s[i])) % prime
	}

	return result
}
