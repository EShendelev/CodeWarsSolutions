package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	str := ReverseWords("The quick brown fox jumps over the lazy dog.")
// 	fmt.Println(str)
// }

func main() {
	str := "stressed   desserts"

	arrStr := strings.Fields(str)
	fmt.Printf("%#+v\n", arrStr)
	for i, v := range arrStr {
		word := strings.SplitAfter(v, " ")
		word = reverseStr(word)
		arrStr[i] = strings.Join(word, "")
		fmt.Printf("%#+v\n", v)
	}
	res := strings.Join(arrStr, " ")
	fmt.Printf("%#+v\n", res)

	//strings.ReplaceAll(resStr, "_", " ")
	// return resStr // reverse those words
}

func reverseStr(str []string) []string {
	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-1-i] = str[len(str)-1-i], str[i]
	}
	return str
}
