package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	prime := true
	limit := int(math.Sqrt(float64(num)))
	for i := 2; i <= limit; i++ {
		if num%i == 0 {
			prime = false
		}
	}
	return prime
}

func gen(num int) chan int {
	out := make(chan int)
	limit := int(math.Sqrt(float64(num)))
	go func() {
		for i := limit; i > 2; i-- {
			if num%i == 0 {
				out <- i
			}
		}
		close(out)
	}()
	return out
}

func main() {
	num := 600851475143
	c := gen(num)
	for n := range c {
		if isPrime(n) {
			fmt.Println(n)
			break
		}
	}
}

// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?
