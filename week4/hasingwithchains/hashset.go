package main

import "fmt"

type chainItem struct {
	key  string
	next *chainItem
}

type HashSet struct {
	numChains int
	chains    []*chainItem
}

func (hs *HashSet) Add(s string) {
	index := hs.hash(s)
	chain := hs.chains[index]

	for ; chain != nil; chain = chain.next {
		if chain.key == s {
			// do nothing if the key already exists
			return
		}
	}

	// apend the key to the beginning of the cgain
	item := chainItem{key: s, next: hs.chains[index]}
	hs.chains[index] = &item
}

func (hs *HashSet) Del(s string) {
	index := hs.hash(s)
	chain := hs.chains[index]
	var prev *chainItem

	for chain != nil {
		if chain.key == s {
			break
		}
		prev = chain
		chain = chain.next
	}

	if chain != nil {
		if prev != nil {
			prev.next = chain.next
		} else {
			hs.chains[index] = chain.next
		}
	}
}

func (hs *HashSet) Find(s string) bool {
	index := hs.hash(s)
	chain := hs.chains[index]

	for ; chain != nil; chain = chain.next {
		if chain.key == s {
			return true
		}
	}
	return false

}

func (hs *HashSet) Check(chainNum int) {

	chain := hs.chains[chainNum]
	for ; chain != nil; chain = chain.next {
		fmt.Printf("%s ", chain.key)
	}
	fmt.Println("")
}

func (hs *HashSet) hash(s string) int {

	const multiplier = 263
	const prime = 1000000007

	result := uint64(0)
	for i := len(s) - 1; i >= 0; i-- {
		result = ((result * multiplier) + uint64(s[i])) % prime
	}

	return int(result % uint64(hs.numChains))
}

func NewHashSet(numChains int) *HashSet {
	chains := make([]*chainItem, numChains)
	return &HashSet{chains: chains, numChains: numChains}
}
