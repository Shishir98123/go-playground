package main

import "fmt"

func main(){
	// map[keyType]valueType
	ages := map[string]int{
		"shishir" : 22,
		"john" : 34,
	}

	fmt.Println(ages["shishir"], len(ages))


	// make(map[K]V)

	var scores map[string]int   //null map
	fmt.Println(scores, scores["a"])

	scores = make(map[string]int)
	scores["math"] = 90
	fmt.Println(scores, scores["math"])


	users := map[string]string{
		"u1" : "sangam",
		"u2" : "shishir",
		"u3" : "sital",
	}
	fmt.Println(users)
	delete(users, "u2")

	fmt.Println(users)

}