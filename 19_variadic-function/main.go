package main

import "fmt"

// A variadic function in Go is a function that can accept a variable number of arguments for its final parameter, which is declared using an ellipsis (...) before the parameter's type
func sumAll(nums ...int) int {
	total := 0

	for _, currentValue := range nums{
		total = total + currentValue
	}
	fmt.Println(total + 5)
	return total
} 

func main(){
	totalSum := sumAll(1,2,3,4,5)
	fmt.Println(totalSum)

	// anonymous function
	res := func(n int) int {
		return n * 2
	}

	fmt.Println(res(2))

	// IIFE
	res1 := func(a int, b int)int{
		return a+b
	}(7, 19)

	fmt.Println(res1)
}