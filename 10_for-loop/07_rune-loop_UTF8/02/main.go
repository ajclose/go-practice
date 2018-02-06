package main

import "fmt"

func main() {
	for i := 50; i <= 150; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}
