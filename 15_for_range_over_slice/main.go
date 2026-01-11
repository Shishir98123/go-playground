package main

import "fmt"

func main(){
	views := []int{10,20,30,40,50,60}
	total := 0
	//  classic for loop
	for i:=0; i <= len(views)-1 ; i++{
		total = total + views[i]
	}
	fmt.Println(total)

	// using for range
	totalV := 0
	for index, value := range views{
		fmt.Println("day", index, "views", value)
		totalV = totalV + value
	}
	fmt.Println(totalV)
}