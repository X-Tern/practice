package main

import ("fmt";
	"work/palindrome"
)



func main() {
	s := "é😊"
	runes := []rune(s) // convert string to runes

	
}



func main() {


	fmt.Println("Bytes length:", len(s))
	fmt.Println("Character length:", len(runes))
	
	fmt.Println("working golang command")
	fmt.Println(palindrome.IsPalindrome("madam"))
	fmt.Println(palindrome.IsPalindrome("A man, a plan, a canal, Panama"))
}