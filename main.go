package main

import "fmt"
import (
	"geuler/euler"
	"geuler/euler/benchmark"
)

type eulerFunc func() (string, int, error)

func runEulerFunc(ix int, fn eulerFunc) {
	bm := benchmark.Start()
	name, result, err := fn()
	µs := benchmark.End(bm)
	if err != nil {
		fmt.Printf("%04d: %s - %s (%dµs)\n", ix, name, err, µs)
	} else {
		fmt.Printf("%04d: %s - Answer: %d (%dµs)\n", ix, name, result, µs)
	}

}

func main() {
	problems := []eulerFunc{
		euler.Problem0001MultiplesOfThreeAndFive,
		euler.Problem0002FibonacciSequence,
		euler.Problem0003PrimeFactors,
		euler.Problem0004PalindromeProduct,
	}

	for ix, fn := range problems {
		runEulerFunc(ix+1, fn)
	}
}
