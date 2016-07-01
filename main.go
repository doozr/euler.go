package main

import "fmt"
import (
	"github.com/doozr/geuler/euler"
	"github.com/doozr/geuler/euler/benchmark"
)

type eulerFunc func() (string, int, error)

type Result struct {
	index int
	name  string
	value int
	err   error
	µs    int
}

func runEulerFunc(ix int, fn eulerFunc) <-chan Result {
	ch := make(chan Result)
	go func() {
		bm := benchmark.Start()
		name, value, err := fn()
		µs := benchmark.End(bm)
		ch <- Result{ix, name, value, err, µs}
		close(ch)
	}()
	return ch
}

func aggregate(agg chan Result, ch <-chan Result) {
	for result := range ch {
		agg <- result
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
		euler.Problem0008LargestProductInSeries,
		euler.Problem0009SpecialPythagoreanTriple,
		euler.Problem0010SummationOfPrimes,
	}

	aggregateChannel := make(chan Result)

	bm := benchmark.Start()
	for ix, fn := range problems {
		go aggregate(aggregateChannel, runEulerFunc(ix+1, fn))
	}

	for ix := 0; ix < len(problems); ix++ {
		result := <-aggregateChannel
		if result.err != nil {
			fmt.Printf("%04d: %s - %s (%dµs)\n", result.index, result.name, result.err, result.µs)
		} else {
			fmt.Printf("%04d: %s - Answer: %d (%dµs)\n", result.index, result.name, result.value, result.µs)
		}
	}

	fmt.Printf("Total time: %dµs", benchmark.End(bm))
}
