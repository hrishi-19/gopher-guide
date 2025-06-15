package main

import "fmt"

type Point struct {
	X int
	Y int
}

func structDemo() {
	fmt.Println(Point{1, 2})
	v := Point{1, 2}
	v.X = 4
	fmt.Println(v.X)
	z := Point{1, 2}
	p := &z
	p.X = 1e9
	fmt.Println(z)
}
