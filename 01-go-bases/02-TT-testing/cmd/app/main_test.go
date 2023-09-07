package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaxSalaryA(t *testing.T) {

	salaryA := 49999.9
	expectedResult := 0

	obtainedResult := calculateSalaryTax(salaryA)

	assert.Equal(t, expectedResult, obtainedResult, "test failed: the expected result was not obtained")
}

func TestTaxSalaryB(t *testing.T) {

	salaryB := 50000.0
	expectedResult := 17

	obtainedResult := calculateSalaryTax(salaryB)

	assert.Equal(t, expectedResult, obtainedResult, "test failed: the expected result was not obtained")
}

func TestTaxSalaryC(t *testing.T) {

	salaryC := 150000.0
	expectedResult := 27

	obtainedResult := calculateSalaryTax(salaryC)

	assert.Equal(t, expectedResult, obtainedResult, "test failed: the expected result was not obtained")
}
