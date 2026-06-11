//go:build ignore

package main

import "fmt"

// isPalindrome reports whether s is a palindrome considering only alphanumeric
// runes and ignoring case.
func isPalindrome(s string) bool {
	// TODO: implement
	return false
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // expect true
	fmt.Println(isPalindrome("race a car"))                     // expect false
	fmt.Println(isPalindrome(""))                               // expect true
}
