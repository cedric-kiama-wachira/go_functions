package main

import (
	"fmt"
)

func addnumbers(a int, b int) int {
	sum := a + b
	return sum

}

func main() {
	sumOfNumbers := addnumbers(2, 3)
	fmt.Println(sumOfNumbers)
}
