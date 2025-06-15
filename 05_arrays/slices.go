package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func sliceDemo() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var a []int = primes[1:4]
	fmt.Println(a)
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	c := names[0:2]
	b := names[1:3]
	fmt.Println(c, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13} // slice literal
	fmt.Println(q)

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

}
