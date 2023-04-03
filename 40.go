package main

import "fmt"

func printNumbersFromAToBRepeatedly(A int, B int) {
    for i := A; i <= B; i++ {
        for j := 1; j <= i-A+1; j++ {
            fmt.Printf("%d ", i)
        }
        fmt.Println()
    }
}
