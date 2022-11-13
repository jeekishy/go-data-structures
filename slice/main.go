package main

import (
	"fmt"
	"time"
)

func main() {
	var mySlice []string

	// adding value to a slice
	mySlice = append(mySlice, "First element")  // O(1)
	mySlice = append(mySlice, "Second element") // O(1)
	mySlice = append(mySlice, "Third element")  // O(1)
	mySlice = append(mySlice, "Fourth element") // O(1)
	//fmt.Println(mySlice)

	// copying a slice using copy - O(n)
	sliceCopyTime := time.Now()
	sliceCopy := make([]string, len(mySlice)) // make an empty slice with same length as the one you want to copy
	copy(sliceCopy, mySlice)
	fmt.Println(time.Since(sliceCopyTime))
	//fmt.Println(sliceCopy)

	// copy slice using append - O(n)
	extraCopyTime := time.Now()
	extraCopy := make([]string, 0)
	extraCopy = append(extraCopy, mySlice[:]...)
	fmt.Println(time.Since(extraCopyTime))
	//fmt.Println(extraCopy)

	// deleting value from Slice by index
	indexToDelete := 2                                                      // Third element
	mySlice = append(mySlice[:indexToDelete], mySlice[indexToDelete+1:]...) // O(1)
	//fmt.Println(mySlice)

	// deleting a value from a slice by value
	// O(n)
	valueToDelete := "Second element"
	for i, value := range mySlice {
		if value == valueToDelete {
			mySlice = append(mySlice[:i], mySlice[i+1:]...)
			break
		}
	}
	//fmt.Println(mySlice)
}
