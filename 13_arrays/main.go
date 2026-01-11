package main
import "fmt"

// In go, an array is a fixed-size, ordered sequence of elements of the same type.
func main(){
	var marks [3]int

	marks[0] = 10
	marks[1] = 20
	marks[2] = 30

	fmt.Println(marks)

	// array literal

	res := [5]int{2,3,4,5,6}
	fmt.Println("length:",len(res))

	var numbers [3]int = [3]int{3,4,5}

	var fruits = [5]string{"mango", "apple", "banana", "grapes", "peach"}

	fmt.Println(numbers[1])
	fmt.Println(fruits)


}