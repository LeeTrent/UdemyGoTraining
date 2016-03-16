package main
import "fmt"
func main() {
    for ii := 0; ii <= 100; ii++ {
        if ii % 2 == 0 {
            fmt.Println(ii)    
        }
    }
}