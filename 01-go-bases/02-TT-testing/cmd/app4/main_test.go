package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperation(t *testing.T) {
	testCase := []struct {
		name           string
		funcion        string
		ratings        []int
		expectedResult float64
	}{
		{"Minimum", "minimum", []int{9, 10, 3, 6, 8, 9, 7, 7}, 3},
		{"Average", "average", []int{9, 10, 3, 6, 8, 9, 7, 7}, 7.375},
		{"Maximum", "maximum", []int{9, 10, 3, 6, 8, 9, 7, 7}, 10},
	}

	for _, test := range testCase {
		t.Run(test.name, func(t *testing.T) {
			result, err := operation(test.funcion, test.ratings)

			assert.Equal(t, test.expectedResult, result, err)
		})
	}

}

func TestMinumumFunc(t *testing.T) {

	ratings := []int{9, 10, 3, 6, 8, 9, 7, 7}

	expectedResult := 3

	obtainedResult := minFunc(ratings)

	assert.Equal(t, expectedResult, obtainedResult, "test failed: the expected result was not obtained")

}

func TestAverageFunc(t *testing.T) {

	ratings := []int{9, 10, 3, 6, 8, 9, 7, 7}

	expectedResult := 7.375

	obtainedResult := averageFunc(ratings)

	assert.Equal(t, expectedResult, obtainedResult, "test failed: the expected result was not obtained")

}

func TestMaximumFunc(t *testing.T) {

	ratings := []int{9, 10, 3, 6, 8, 9, 7, 7}

	expectedResult := 10

	obtainedResult := maxFunc(ratings)

	assert.Equal(t, expectedResult, obtainedResult, "test failed: the expected result was not obtained")

}
