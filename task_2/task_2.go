package main

import (
	"fmt"
	"unicode"
)

func word_count(word string) map[string]int {
	my_map := map[string]int{}
	for _, char := range word {
		charStr := string(char)
		if count, exists := my_map[charStr]; exists {
			my_map[charStr] = count + 1
		} else {
			my_map[charStr] = 1
		}
	}
	return my_map
}

func palindrome_check(word string) bool {
	var cleanChars []rune
	for _, char := range word {
		if unicode.IsLetter(char) {
			cleanChars = append(cleanChars, unicode.ToLower(char))
		}
	}

	left_pointer := 0
	right_pointer := len(cleanChars) - 1

	for left_pointer < right_pointer {
		if cleanChars[left_pointer] != cleanChars[right_pointer] {
			return false
		}
		left_pointer++
		right_pointer--
	}
	return true
}

func main() {
	result := word_count("hello")
	fmt.Println(result)
	fmt.Println(palindrome_check("A man, a plan, a canal: Panama"))

}
