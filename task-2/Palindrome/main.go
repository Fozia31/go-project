package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.TrimSpace(s)
	runes := []rune(strings.ToLower(s)) // convert to runes

	left, right := 0, len(runes)-1
	for left < right {

		for left < right && !unicode.IsLetter(runes[left]) && !unicode.IsDigit(runes[left]) {
			left++
		}

		for left < right && !unicode.IsLetter(runes[right]) && !unicode.IsDigit(runes[right]) {
			right--
		}

		if runes[left] != runes[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	st, _ := reader.ReadString('\n')
	st = strings.TrimSpace(st) // remove newline

	fmt.Println(isPalindrome(st))
}
