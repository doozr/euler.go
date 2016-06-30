package euler

/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?


Answer:
    232792560
*/

func makeDivisible(num, divisor int) int {
	for x := 1; x < divisor; x++ {
		if num * x % divisor == 0 {
			return x
		}
	}
	return divisor
}

func smallestMultiple(limit int) int {
	result := 1
	for divisor := 2; divisor <= limit; divisor++ {
		result *= makeDivisible(result, divisor)
	}
	return result
}

func Problem0005SmallestMultiple() (string, int, error) {
	name := "Smallest multiple"
	expected := 232792560
	actual := smallestMultiple(20)
	return name, actual, AssertEqual(expected, actual)
}
