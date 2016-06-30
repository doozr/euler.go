package main

import "fmt"
import "geuler/euler"

type eulerFunc func() (string, int, error)

func runEulerFunc(ix int, fn eulerFunc) {
	name, result, err := fn()
	if err != nil {
		fmt.Printf("%04d: %s - %s\n", ix, name, err)
	} else {
		fmt.Printf("%04d: %s - Answer: %d\n", ix, name, result)
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
