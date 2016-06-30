package main

import "fmt"

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
		Problem0001MultiplesOfThreeAndFive,
		Problem0002FibonacciSequence,
		Problem0003PrimeFactors,
	}

	for ix, fn := range problems {
		runEulerFunc(ix+1, fn)
	}
}
