package main

import "fmt"

func main() {

	salary := 70000

	if salary < 150000 {
		err := SalaryError{"error: the salary entered does not reach the taxable minimum"}
		fmt.Println(err.Error())
	} else {
		fmt.Println("must pay tax")
	}

}

type SalaryError struct {
	err string
}

func (se SalaryError) Error() string {
	return se.err
}
