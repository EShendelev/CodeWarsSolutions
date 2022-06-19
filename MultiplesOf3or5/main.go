package main

import "fmt"

func main() {
	fmt.Println(Multiple3And5(10))
}

func Multiple3And5(number int) int {
	sum := 0
	for i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		} else if number < 0 {
			sum = -1
		}
	}
	return sum
}
