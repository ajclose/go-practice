package main

import "fmt"

func greatest(x ...float64) float64 {
	var greatest float64
	for i, v := range x {
		if v > greatest || i == 0 {
			greatest = v
		}
	}
	return greatest
}

func main() {
	fmt.Println(greatest(-9, -2, -3, -6))
}
