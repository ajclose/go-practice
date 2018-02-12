package main

import "fmt"

func main() {
	num := 20
	for i := num * num; i > 0; i += num {
		divisible := true
		for j := 2; j <= num; j++ {
			if i%j != 0 {
				divisible = false
				break
			}
		}
		if divisible {
			fmt.Println(i)
			break
		}
	}
}

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
