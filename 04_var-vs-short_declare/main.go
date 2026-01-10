package main

import (
	"fmt"
)

func main(){
	var firstName string = "Shishir"
	var lastName = "Khatiwada"  // inferred to string
	
	fmt.Println("My name is",firstName, lastName)

	// short form " := "

	subscribers := 50000  // we don't need to mention its type it will automatically infer.

	subscribers = subscribers + 10000

	likes , comments := 20000, 500

	fmt.Println("Subscribers:", subscribers)
	fmt.Println("likes:", likes)
	fmt.Println("comments:", comments)
}
