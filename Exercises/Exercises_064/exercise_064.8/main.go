package main

import "fmt"

func main () {
    fmt.Println(process(54))
}

func process(val int) (float64, bool) {
    return float64(val / 2), (val % 2 == 0)
}