package euler

import "github.com/doozr/geuler/euler/assert"

/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.


Answer:
    233168
*/

// MultiplesOfThreeAndFive calculate the sum of multiples of 3 and 5 under limit
func MultiplesOfThreeAndFive(limit int) int {
	total := 0
	for ix := 1; ix < limit; ix++ {
		if ix%3 == 0 || ix%5 == 0 {
			total += ix
		}
	}
	return total
}

// Problem0001MultiplesOfThreeAndFive calculate the sum of multiples of 3 and 5 under 1000
func Problem0001MultiplesOfThreeAndFive() (string, int, error) {
	name := "Multiples of three and five"
	expected := 233168
	actual := MultiplesOfThreeAndFive(1000)
	return name, actual, assert.Equal(expected, actual)
}
