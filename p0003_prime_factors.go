package main

import "github.com/doozr/geuler/euler/math"

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
	total := highestPrimeFactor(600851475143)
	expected := 6857
	if total != expected {
		return name, total, EulerError{expected, total}
	}
	return name, total, nil
}
