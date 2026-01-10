package main

import (
	"fmt"
	"strings"
)

func main(){
	firstName := "Shishir"
	lastName := "Khatiwada"
	fullName := firstName + " " + lastName

	fmt.Println(fullName)
	fmt.Println(strings.ToUpper(fullName))
}