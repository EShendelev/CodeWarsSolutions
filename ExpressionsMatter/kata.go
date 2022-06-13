package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(ExpressionMatter(1, 7, 1))
}

func ExpressionMatter(a int, b int, c int) int {
	s := []int{a * b * c, a * (b + c), a + b*c, (a + b) * c, a*b + c, a + b + c}
	sort.Ints(s)
	fmt.Printf("%#+v\n", s)
	return s[len(s)-1]
}
