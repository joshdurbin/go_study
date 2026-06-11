//go:build ignore

package main

import (
	"fmt"
	"unicode"
)

// isPalindrome uses opposite-ends pointers, skipping non-alphanumeric runes.
// O(n) time, O(n) space for the rune slice (could be O(1) by using utf8.DecodeRune).
func isPalindrome(s string) bool {
	r := []rune(s)
	i, j := 0, len(r)-1
	for i < j {
		switch {
		case !isAlnum(r[i]):
			i++
		case !isAlnum(r[j]):
			j--
		case unicode.ToLower(r[i]) != unicode.ToLower(r[j]):
			return false
		default:
			i++
			j--
		}
	}
	return true
}

func isAlnum(r rune) bool { return unicode.IsLetter(r) || unicode.IsDigit(r) }

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car"))                     // false
	fmt.Println(isPalindrome(""))                               // true
}
