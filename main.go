package main

import "fmt"

func sum(x int, y int) int{
	z :=x+y
	return z
}

func main() {
	fmt.Printf("Hi theree ..\n")
	r := sum(2,3)
	fmt.Printf("sum of 2 and 3 is %v", r)
}
