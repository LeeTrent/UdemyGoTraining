package main

import "fmt"

func main() {
    var userInput string
    fmt.Print("Please enter your name: ")
    fmt.Scan(&userInput)
    fmt.Println("Hello", userInput)
}