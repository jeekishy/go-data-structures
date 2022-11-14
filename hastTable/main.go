package main

import (
	"fmt"
	"hash/fnv"
)

// implement a simple custom hash table

// HastTable struct
type HastTable struct {
	size int
	data map[int]int
}

func main() {
	ht := NewHashTable()

	// adding to a hash table
	ht.set("apple", 10000)     // O(1)
	ht.set("pineapple", 20000) // O(1)
	ht.set("smith", 30000)     // O(1)

	// getting value from hash table
	fmt.Println(ht.get("apple"))     // O(1)
	fmt.Println(ht.get("pineapple")) // O(1)
	fmt.Println(ht.get("smith"))     // O(1)

	// deleting value from hash table
	ht.delete("pineapple")           // O(1)
	fmt.Println(ht.get("pineapple")) // O(1)
}

// NewHashTable get new hashTable
func NewHashTable() *HastTable {
	return &HastTable{
		size: 0,
		data: make(map[int]int),
	}
}

// Set adds a value to the hash table
func (h *HastTable) set(key string, value int) {
	h.data[hash(key)] = value
	h.size++
}

// get retrieves a value from hash table
func (h *HastTable) get(key string) int {
	return h.data[hash(key)]
}

// delete removes data from hast table
func (h *HastTable) delete(key string) bool {
	// check map has key
	if _, ok := h.data[hash(key)]; ok {
		delete(h.data, hash(key))
		return true
	}
	return false
}

// hash generates an int from provide key string
func hash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32())
}
