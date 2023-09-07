package main

import "fmt"

/* Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
1. Categoría C, su salario es de $1.000 por hora.
2. Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
3. Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes, la categoría y que devuelva su salario. */

func main() {
	fmt.Println(calculateSalary("A", 10560))
}

// minutos y categoria
func calculateSalary(category string, minutes float64) float64 {
	// inicializo el salario y el precio por minuto que va a calcular el precio por 1 hora (depende la categoría) dividio en 60 minutos
	var salary float64
	var priceMinute float64
	switch category {
	case "C":
		priceMinute = 1000.0 / 60.0
		salary = minutes * priceMinute
	case "B":
		priceMinute = 1500.0 / 60.0
		sum := minutes * priceMinute
		salary = sum * 1.2
	case "A":
		priceMinute = 3000.0 / 60.0
		sum := minutes * priceMinute
		salary = sum * 1.5
	}
	return salary
}
