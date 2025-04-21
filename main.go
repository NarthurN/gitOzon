//package to commit
package main

import "fmt"

func main() {
	sum56 := sum(5, 6)
	fmt.Println(sum56)
}

func sum(a, b int) int {
	return a + b
}

func anotherSum(a, b int) int {
	return 5 - 2
}