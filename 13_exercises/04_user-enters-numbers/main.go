package main

import "fmt"

func main() {
	var largeNumber int
	var smallNumber int
	fmt.Println("Enter a large number")
	fmt.Scan(&largeNumber)
	fmt.Println("Enter a small number")
	fmt.Scan(&smallNumber)
	fmt.Println(largeNumber, "%", smallNumber, " = ", largeNumber%smallNumber)
}
