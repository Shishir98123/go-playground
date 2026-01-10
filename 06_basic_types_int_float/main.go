package main

import "fmt"

func main(){
	fmt.Println("hello")
	views1 := 1000
	views2 := 2000

	totalViews := views1 + views2

	likes := 10
	likes++
	likes++

	avgViews := totalViews/2

	fmt.Println(totalViews)
	fmt.Println(likes)
	fmt.Println(avgViews)

	rating1 := 4.5
	rating2 := 9.8

	avgRating := (rating1 + rating2)/2

	fmt.Println(avgRating)
	
}
