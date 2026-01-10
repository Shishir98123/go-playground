package main

import "fmt"

func main(){
	isLogged := true;
	isAdmin := false;
	hasSubscription := true;
	
	canOpenDashboard := isLogged && hasSubscription

	canDeleteDashboard := isAdmin || (isLogged && hasSubscription)

	fmt.Println(canOpenDashboard)
	fmt.Println(canDeleteDashboard)

	age := 45
	isAdult := age > 18 && age < 45
	fmt.Println(isAdult)
}