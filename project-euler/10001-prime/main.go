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

func main() {
	primes := []int{}
	length := 10001
	for i := 2; i > 0; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
		if len(primes) == length {
			fmt.Println(i)
			break
		}
	}
}
