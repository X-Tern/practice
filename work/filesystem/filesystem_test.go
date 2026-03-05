package filesystem


import (
	"testing"
)


func TestIsPalindromeTrue(t *testing.T) {
	if !IsPalindrome("madam") {
		t.Error("madam should be true")
	}
}

func TestIsPalindromeFalse(t *testing.T) {
	if IsPalindrome("hello") {
		t.Error("hello should be false")
	}
}