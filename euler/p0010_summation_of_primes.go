package euler

import (
	"github.com/doozr/geuler/euler/assert"
	"github.com/doozr/geuler/euler/math"
)

/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.


Answer:
    142913828922
 */

func summationOfPrimes(limit int) int {
	total := 0
	primes := math.Primes(2000000)
	for i := 0; i < limit; i ++ {
		total += <- primes
	}
	return total
}

func Problem0010SummationOfPrimes() (string, int, error) {
	name := "Summation of primes"
	expected := 142913828922
	actual := summationOfPrimes(2000000)
	return name, actual, assert.Equal(expected, actual)
}