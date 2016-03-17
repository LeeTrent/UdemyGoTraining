package main

import "fmt"

func main () {   
    process := func (val int) (float64, bool) {
        return float64(val / 2), (val % 2 == 0)
    }
    fmt.Println(process(54))
}