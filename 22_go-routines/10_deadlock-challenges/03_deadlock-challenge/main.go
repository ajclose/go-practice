package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

// Why does this only print zero?
// And what can you do to get it to print all 0 - 9 numbers?
