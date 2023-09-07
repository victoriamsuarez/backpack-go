package main

import (
	"fmt"
)

func main() {

	salary := 150_000

	if salary < 150_000 {
		errorMessage := SalaryError(salary)
		fmt.Println(errorMessage)
	} else {
		fmt.Println("ok")
	}

}

func SalaryError(salary int) string {
	var ErrMinimumTaxable = fmt.Errorf("error: the taxable minimum is 150,000 and the salary entered is: %d", salary)
	return ErrMinimumTaxable.Error()
}
