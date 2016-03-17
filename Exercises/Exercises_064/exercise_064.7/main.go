package main

import "fmt"

func main() {
	var sum int = 0
	for ii := 0; ii < 1000; ii++ {
		if ii%3 == 0 || ii%5 == 0 {
			sum += ii
		}
	}
	fmt.Println(sum)
}
