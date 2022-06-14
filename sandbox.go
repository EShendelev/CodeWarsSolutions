package main

import "fmt"

func main() {
	s := "love"
	count := 0
	for _, i := range s {
		count += int(i) - 'a' + 1
		fmt.Printf("%+v, %+v, %+v \n", i, int(i), 'a')
	}
	fmt.Println(count)
}
