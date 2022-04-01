package main

import "fmt"

//Declare factorial function here
func factorial(number int) int {
	if number == 0 {
		return 1
	} else {
		return number * factorial(number-1)
	}
}

func main() {
	//Insert Code here
	fmt.Println("Input number")
	var userInput int
	fmt.Scan(&userInput)
	fmt.Println(factorial(userInput))
}
