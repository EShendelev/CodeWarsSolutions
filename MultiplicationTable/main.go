package main

import "fmt"

func main() {
	fmt.Println(MultiplicationTable(4))
}

func MultiplicationTable(size int) [][]int {
	res := make([][]int, size)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, size)
		for j := 0; j < len(res); j++ {
			res[i][j] = (i + 1) * (j + 1)
		}
	}
	return res
}
