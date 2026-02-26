package palindrome

import "strings"


func isPalindrome(input string) bool {

	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")

	runes := []rune(input)

	left := 0
	right := len(runes)

	for left < right {
		if runes[left] != runes[right]{
			return False
		}
	left ++
	right --
	}
	return true
}

fmt.Println(isPalindrome("Terna"))