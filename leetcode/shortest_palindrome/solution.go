package main

import "fmt"

/*
	idea:
	    doing shifting then substring checking for original string and reversed string until found palindrome
		return number of shifted word from reversed string + original string

		example 1:
		a b c d
		d c b a

		--- shift ---

		. . . a b c d  ->  a b c d  ->  a b c  ->  a b  ->  a
		d c b a . . .  ->  d c b a  ->  c b a  ->  b a  ->  a

		================================================

		example 2:

		a a c e c a a a
		a a a c e c a a

		--- shift ---

		. a a c e c a a a   ->  a a c e c a a a  ->  a a c e c a a
		a a a c e c a a .   ->  a a a c e c a a  ->  a a c e c a a
*/

// reverse function to reverse string
func reverse(s string) (reversed string) {
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}

	return
}

// shortestPalindrome get shortest palindrome from given string
func shortestPalindrome(s string) string {
	// define original string and reversed string
	reversed := reverse(s)
	s1 := s
	s2 := reversed
	shiftCount := 0

	// loop until string is last 1 char
	for {
		// check if palindrome, then return additional char + original string
		if s1 == s2 {
			break
		}

		// shifting
		s1 = s1[0 : len(s1)-1]
		s2 = s2[1:]
		shiftCount++
	}

	return reversed[0:shiftCount] + s
}

func main() {
	fmt.Println(shortestPalindrome("abcd"))
	fmt.Println(shortestPalindrome("aacecaaa"))
}
