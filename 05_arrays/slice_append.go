package main

import (
	"fmt"
)

func aliceAppend() {
	var s []int
	printArraySlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printArraySlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printArraySlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printArraySlice(s)
}

func printArraySlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
