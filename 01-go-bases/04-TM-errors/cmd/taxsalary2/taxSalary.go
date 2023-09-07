package main

import (
	"errors"
	"fmt"
)

type SalaryError struct {
	msg string
}

func (se SalaryError) Error() string {
	return se.msg
}

func main() {

	salary := 7000

	if salary <= 10000 {
		err := SalaryError{"error: salary is less than 10000"}
		if errors.Is(err, SalaryError{
			msg: "error: salary is less than 10000",
		}) {
			fmt.Println(err.Error())
			return
		}
	}
	fmt.Println("must pay")
}
