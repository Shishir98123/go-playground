package main

import "fmt"

func main(){
	points := map[string]int{
		"a" : 10,
		"b" : 0,  // valid value
	}

	fmt.Println("a", points["a"])
	fmt.Println("b", points["b"])
	fmt.Println("c", points["c"])

	valB, okB := points["b"]
	fmt.Println(valB, okB)

	valC, okC := points["c"]
	fmt.Println(valC, okC)


	// if else with statement
	if val, ok := points["c"]; ok{
		fmt.Println(val)
	}else{
		fmt.Println("not present")
	}

	// for range in map
	prices := map[string]int{
		"xyz":399,
		"abc": 599,
	}
	total := 0
	for items, price := range prices{
		fmt.Println(items, price)
		total += price
	}
	fmt.Println(total)

	for item := range prices{
		fmt.Println(item)
	}
}