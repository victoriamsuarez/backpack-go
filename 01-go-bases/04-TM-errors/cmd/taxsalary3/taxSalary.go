package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := SalaryError()

	salary := 70000

	if salary <= 10000 {
		if errors.Is(err, ErrSalaryLow) {
			fmt.Println(err)
		}
	} else {
		fmt.Println("must pay", result)
	}

}

var ErrSalaryLow = errors.New("error: salary is less than 10000")

func SalaryError() (string, error) {
	return "", ErrSalaryLow
}
