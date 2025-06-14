package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	ar i int = 42
    var f float64 = 3.14
    var b bool = true
    var s string = "Hello, Go!"

    fmt.Printf("Integer: %d\n", i)
    fmt.Printf("Float: %.2f\n", f)
    fmt.Printf("Boolean: %t\n", b)
    fmt.Printf("String: %s\n", s)
	
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}