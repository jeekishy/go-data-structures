package main

import (
	"fmt"
	"hash/fnv"
)

// implement a simple custom hash table

// HastTable struc
type HastTable struct {
	size int
	data map[int]int
}

func main() {
	ht := NewHashTable()
	ht.set("apple", 10000)
	ht.set("pineapple", 20000)
	ht.set("smith", 30000)

	fmt.Println(ht.get("apple"))
	fmt.Println(ht.get("pineapple"))
	fmt.Println(ht.get("smith"))
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

// hash generates a int from provide key string
func hash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32())
}
