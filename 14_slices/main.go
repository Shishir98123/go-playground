package main
import "fmt"
func main(){
	// common collection type
	// dynamic and can grow
	// can be resized

	// A slice is a descriptor that points to a contiguous segment of an underlying array.

	results  := []string{"shishir", "sangam", "sital"}
	// this creates:
	// 1) An array holding the data
	// 2) A slice pointing to that array
	fmt.Println(results)

	// DECLARATION

	var s []int // s == nill,  len(s) == 0,  cap(s) == 0
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s)) // capacity

	// LITERAL
	p := []int{1,2,3}
	fmt.Println(p)

	// USING MAKE
	a := make([]int, 5)   // length = 5, capacity = 5
	b := make([]string, 5, 10)   // length = 5, capacity = 10
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(len(a))
	fmt.Println(cap(a)) 


	nums := []int{1,2,3,4,5,6}
	fmt.Println(len(nums))
	fmt.Println(cap(nums))

	nums1 := nums[2:5]
	fmt.Println(nums1)
	fmt.Println(len(nums1))
	fmt.Println(cap(nums1))

	// slicing syntax
	// [low : high : max]
	d := []int{1,2,3,4,5}
	e := d[1:3:3]
	fmt.Println(len(e))
	fmt.Println(cap(e))
}