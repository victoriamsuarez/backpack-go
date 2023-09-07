package main

import "testing"

func TestCalculateAvg(t *testing.T) {

	ratings := []int{8, 7, 9, 10, 7, 9, 9, 8}
	expectedResult := 8.375

	obtainedResult := calculateAverage(ratings)

	if obtainedResult != expectedResult {
		t.Fatal("the result was expected", expectedResult, "the result was obtained", obtainedResult)
	}
}
