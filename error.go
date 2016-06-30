package main

import "fmt"

type EulerError struct {
	expected int
	actual   int
}

func (e EulerError) Error() string {
	return fmt.Sprintf("Expected %d, actual %d", e.expected, e.actual)
}
