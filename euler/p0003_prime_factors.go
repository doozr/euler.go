package euler

import "geuler/euler/math"

/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?


Answer:
    6857
 */

func highestPrimeFactor(n int) int {
	factors := math.PrimeFactors(n)
	if len(factors) > 0 {
		return factors[len(factors)-1]
	}
	return 0
}

// Problem0003PrimeFactors calculates the highest prime factor of 600851475143
func Problem0003PrimeFactors() (string, int, error) {
	name := "Prime factors"
	expected := 6857
	actual := highestPrimeFactor(600851475143)
	return name, actual, AssertEqual(expected, actual)
}
