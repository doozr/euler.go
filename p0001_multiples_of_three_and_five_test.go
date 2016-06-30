package main

import (
	"testing"
)

/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.


Answer:
    233168
*/

func MultiplesOfThreeAndFive() int {
	total := 0
	for ix := 1; ix < 1000; ix++ {
		if ix%3 == 0 || ix%5 == 0 {
			total += ix
		}
	}
	return total
}

// TestProblem0001MultipleOfThreeAndFive calculates the sum of multiples of 3 and 5 under 1000
func TestProblem0001MultipleOfThreeAndFive(t *testing.T) {
	total := MultiplesOfThreeAndFive()
	expected := 233168
	if total != expected {
		t.Errorf("Expected %d, received %d", expected, total)
	}
}
