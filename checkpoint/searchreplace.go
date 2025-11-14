package main

import "fmt"

func searchreplace(input, search, replace string) string {
	search_rune := []rune(search)
	replace_rune := []rune(replace)
	if len(search_rune) != 1 || len(replace_rune) != 1 {
		return input
	}

	s := search_rune[0]
	r := replace_rune[0]
	results := ""
	for _, char := range input {
		if char == s {
			results += string(r)
		} else {
			results += string(char)
		}
	}
	return results
}

func main() {
	input := "hello"
	s := "h"
	r := "j"
	fmt.Println(searchreplace(input, s, r))
}
