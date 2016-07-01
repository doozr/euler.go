package main

import "fmt"
import (
	"github.com/doozr/geuler/euler"
	"github.com/doozr/geuler/euler/benchmark"
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
		euler.Problem0005SmallestMultiple,
		euler.Problem0006SumSquareDifference,
		euler.Problem0007TenThousandAndFirstPrime,
	}

	bm := benchmark.Start()
	for ix, fn := range problems {
		runEulerFunc(ix+1, fn)
	}
	fmt.Printf("Total time: %dµs", benchmark.End(bm))

}
