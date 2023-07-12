// HELLO WORLD
package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func main() {

	fmt.Println("HELLO WORLD")

	num := 5
	result := factorial(num)
	fmt.Printf("The factorial of %d is %d\n", num, result)

}
