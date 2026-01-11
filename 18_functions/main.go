package main

import "fmt"

func add(x int, y int)int{
	return x+y
}

func sumAndProduct(x, y int) (int, int){
	sum := x+y
	product := x*y
	return sum, product
}

func main(){
	fmt.Println(add(1,3))
	sum , product := sumAndProduct(4,6)
	fmt.Println(sum, product)

}