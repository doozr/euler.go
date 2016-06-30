package euler

import "geuler/euler/assert"

/*
The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385

The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025

Hence the difference between the sum of the squares of the first ten
natural numbers and the square of the sum is 3025 - 385 = 2640.

Find the difference between the sum of the squares of the first one
hundred natural numbers and the square of the sum.


Answer:
    25164150
*/

func sumSquareDifference(limit int) int {
	sumOfSquares := 0
	squareOfSum := 0
	for x := 1; x <= limit; x++ {
		sumOfSquares += x * x
		squareOfSum += x
	}
	squareOfSum *= squareOfSum
	return squareOfSum - sumOfSquares
}

func Problem0006SumSquareDifference() (string, int, error) {
	name := "Sum square difference"
	expected := 25164150
	actual := sumSquareDifference(100)
	return name, actual, assert.Equal(expected, actual)
}
