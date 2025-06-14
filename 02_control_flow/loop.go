package main

import "fmt"

func main() {
    // Basic for loop
    for i := 1; i <= 5; i++ {
        fmt.Println("Count:", i)
    }

    // While-style loop
    n := 1
    for n < 5 {
        fmt.Println("n is", n)
        n++
    }

    // Infinite loop (with break)
    for {
        fmt.Println("Inside infinite loop")
        break
    }
}
