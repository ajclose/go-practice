package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)

	var yourAge int
	yourAge = 28
	fmt.Printf("%T %v \n", yourAge, yourAge)

	// this doesn't work:
	// fmt.Println(myAge + yourAge)

	// this converstion works:
	// fmt.Println(int(myAge) + yourAge)
}
