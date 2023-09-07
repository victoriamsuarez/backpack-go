package main

import (
	"errors"
	"fmt"
)

func main() {

	var hours uint = 86
	price := 2500

	salaryMonth, err := MonthlySalary(hours, price)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Monthly salary: %.2f\n", salaryMonth)
	}

}

var ErrHoursMinWorked = errors.New("error: the worker cannot have worked less than 80 hours per month")

func MonthlySalary(hoursWorked uint, priceHour int) (float64, error) {

	var salaryTotal float64 = float64(hoursWorked) * float64(priceHour)

	if hoursWorked < 80 {
		return 0.0, ErrHoursMinWorked
	}

	if salaryTotal >= 150_000 {
		salaryTotal = salaryTotal * 0.90
	}

	return salaryTotal, nil
}
