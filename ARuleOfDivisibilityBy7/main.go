package main

import "fmt"

func main() {
	fmt.Println(Seven(1603))
}

func Seven(n int64) []int {
	c := 0
	for ; n >= 100; n = n/10 - 2*(n%10) {
		c += 1
	}
	return []int{int(n), c}
}
