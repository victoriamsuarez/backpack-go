package main

import "testing"

func TestCalculateSalaryCategoryA(t *testing.T) {

	category := "A"
	minutesMounth := 480.0
	expectedResult := 36000.0

	obtainedResult := calculateSalary(category, minutesMounth)

	if obtainedResult != expectedResult {
		t.Fatal("the result was expected", expectedResult, "the result was obtained", obtainedResult)
	}

}

func TestCalculateSalaryCategoryB(t *testing.T) {

	category := "B"
	minutesMounth := 480.0
	expectedResult := 14400.0

	obtainedResult := calculateSalary(category, minutesMounth)

	if obtainedResult != expectedResult {
		t.Fatal("the result was expected", expectedResult, "the result was obtained", obtainedResult)
	}

}

func TestCalculateSalaryCategoryC(t *testing.T) {

	category := "C"
	minutesMounth := 450.0
	expectedResult := 7500.000000000001

	obtainedResult := calculateSalary(category, minutesMounth)

	if obtainedResult != expectedResult {
		t.Fatal("the result was expected", expectedResult, "the result was obtained", obtainedResult)
	}

}
