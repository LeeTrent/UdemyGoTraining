package main
import "fmt"
func main() {
    for ii := 1; ii <= 100; ii++ {
        if ii % 3 == 0 && ii % 5 == 0 {
            fmt.Println("FizzBuzz")    
        } else if ii % 3 == 0 {
            fmt.Println("Fizz")
        } else if ii % 5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(ii)
        }
    }
}