package main

import (
	"fmt"
)

/* Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar
el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más
de $150.000 se le descontará además un 10 % (27% en total). */

func main() {
	fmt.Println("The tax on this salary is: ", calculateSalaryTax(30.0), "%")
}

func calculateSalaryTax(salary float64) int {
	tax := 17
	if salary >= 150000.0 {
		return tax + 10
	} else if salary >= 50000.0 && salary < 150000.0 {
		return tax
	} else {
		return 0
	}
}
