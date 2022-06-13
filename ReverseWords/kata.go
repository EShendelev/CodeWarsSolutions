package main

import "strings"

func ReverseWords(str string) string {
	arrStr := strings.Fields(str)
	for i, v := range arrStr {
		word := strings.Split(v, "")
		word = reverseStr(word)
		arrStr[i] = strings.Join(word, "")
	}
	res := strings.Join(arrStr, " ")
	return res
}

func reverseStr(str []string) []string {
	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-1-i] = str[len(str)-1-i], str[i]
	}
	return str
}
