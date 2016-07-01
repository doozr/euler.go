package euler

import (
	"github.com/doozr/geuler/euler/assert"
	"github.com/doozr/geuler/euler/math"
)

/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10001st prime number?


Answer:
    104743
*/

func nthPrime(n int) int {
	ch := make(chan int)
	prime := 0

	go math.Primes(200000, ch)
	for x := 0; x < n; x += 1 {
		prime = <-ch
	}
	return prime
}

func Problem0007TenThousandAndFirstPrime() (string, int, error) {
	name := "10001st prime"
	expected := 104743
	actual := nthPrime(10001)
	return name, actual, assert.Equal(expected, actual)
}
