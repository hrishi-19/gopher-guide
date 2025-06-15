package main

import "fmt"

func arrayDemo() {
	var arr1 [5]string // Array of 5 strings
	arr1[0] = "Hello"
	arr1[1] = "World"
	arr1[2] = "Go"
	arr1[3] = "is"
	arr1[4] = "awesome"
	arr2 := [3]int{1, 2, 3} // Array of 3 integers
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println("Length of arr1:", len(arr1))
	fmt.Println("Length of arr2:", len(arr2))
	fmt.Println("First element of arr1:", arr1[0])
	fmt.Println("Second element of arr2:", arr2[1])
	fmt.Println("Array 1 is of type:", fmt.Sprintf("%T", arr1))
	fmt.Println("Array 2 is of type:", fmt.Sprintf("%T", arr2))
	fmt.Println("Array 1 is empty:", arr1 == [5]string{})
	fmt.Println("Array 2 is empty:", arr2 == [3]int{})
	fmt.Println("Array 1 is equal to another array:", arr1 == [5]string{"Hello", "World", "Go", "is", "awesome"})
	fmt.Println("Array 2 is equal to another array:", arr2 == [3]int{1, 2, 3})
	fmt.Println("Array 1 is not equal to another array:", arr1 != [5]string{"Hello", "Go", "World", "is", "awesome"})
	fmt.Println("Array 2 is not equal to another array:", arr2 != [3]int{1, 3, 2})
}
