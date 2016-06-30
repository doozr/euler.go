package euler

import (
	"math"
	"fmt"
	"geuler/euler/assert"
)

/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 x 99.

Find the largest palindrome made from the product of two 3-digit numbers.


Answer:
906609
*/

func isPalindrome(s string) bool {
	length := len(s)
	halfway := int(length / 2)
	end := length - 1
	for ix := 0; ix <= halfway; ix++ {
		if s[ix] != s[end - ix] {
			return false
		}
	}
	return true
}

func maxPalindromeProduct(start, end int) int {
	lowerBound := start
	product := 0
	for x := end; x > lowerBound; x-- {
		for y := end; y > lowerBound; y-- {
			candidate := x * y
			if isPalindrome(fmt.Sprintf("%d", candidate)) {
				lowerBound = int(math.Min(float64(x), float64(y)))
				if candidate > product {
					product = candidate
				}
			}
		}
	}
	return product
}

func Problem0004PalindromeProduct() (string, int, error) {
	name := "Palindrome product"
	expected := 906609
	actual := maxPalindromeProduct(100, 999)
	return name, actual, assert.Equal(expected, actual)
}