package main

import "fmt"

func main() {
	fmt.Printf("%v\n", Anagrams("laser", []string{"lazing", "lazy", "lacer"}))
}

func Anagrams(word string, words []string) []string {
	if len(words) == 0 {
		return nil
	} else if len(words) == 1 {
		return words
	}
	wordSum := sumOfRune(word)
	res := make([]string, 0)
	findAnagram := false
	for i := 0; i < len(words); i++ {
		if wordSum == sumOfRune(words[i]) {
			findAnagram = true
			res = append(res, words[i])
		}
	}

	if findAnagram {
		return res
	}
	return nil
}

func sumOfRune(word string) int {
	sum := 0

	for i := range word {
		sum += int(word[i])
	}
	return sum
}
