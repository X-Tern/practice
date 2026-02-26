package palindrome

import "strings"


func IsPalindrome(input string) bool {

	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")

	runes := []rune(input)

	left := 0
	right := len(runes) -1

	for left < right {
		if runes[left] != runes[right]{
			return false
		}
	left ++
	right --
	}
	return true
}

