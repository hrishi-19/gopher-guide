package main

import "fmt"

func handlePanic() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()
    panic("Something went wrong")
}
