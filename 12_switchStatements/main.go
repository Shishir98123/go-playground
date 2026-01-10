package main

import "fmt"

func main(){
	day := 6

	switch(day){
		case 1:
			fmt.Println("Sunday")
		case 2:
			fmt.Println("Monday")
		case 3:
			fmt.Println("Tuesday")
		case 4:
			fmt.Println("Wednesday")
		case 5:
			fmt.Println("Thrusday")
		case 6:
			fmt.Println("Friday")
		case 7:
			fmt.Println("Saturday")
		default :
			fmt.Println("Enter number between 1 to 7.")
	}
}