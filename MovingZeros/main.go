package main

import "fmt"

func main() {
	arr := []int{9, 0, 9, 1, 2, 1, 1, 3, 1, 9, 0, 0, 9, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(MoveZeros(arr))
}

func MoveZeros(arr []int) []int {
	// for i, v := range arr {
	// 	if v == 0 {
	// 		arr = append(arr[:i], arr[i+1:]...)
	// 		arr = append(arr, 0)
	// 	}
	// }
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			continue
		}
		arr[j] = arr[i]
		if i != j {
			arr[i] = 0
		}
		j++
	}
	return arr
}
