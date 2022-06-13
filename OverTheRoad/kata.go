package main

import "fmt"

func main() {
	fmt.Println(OverTheRoad(7, 11))
}

func OverTheRoad(address int, n int) int {
	return (n + n + 1) - address
}
