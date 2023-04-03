package main

import "fmt"

func printNumbersFromAToB(A int, B int) {
    for i := A; i <= B; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d ", i)
        }
        fmt.Println()
    }
}
