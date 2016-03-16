package main

import "fmt"

func main() {
    var smallNumber int
    var largeNumber int
    
    fmt.Print("Enter small number: ")
    fmt.Scan(&smallNumber)
    
    fmt.Print("Enter large number: ")
    fmt.Scan(&largeNumber)
    
    if smallNumber != 0 {
        fmt.Println(largeNumber, "/", smallNumber, "=", (largeNumber / smallNumber ) )
        fmt.Println("Remainder is", largeNumber % smallNumber)
    } else {
        fmt.Println("Cannot divide by zero")
    }
}