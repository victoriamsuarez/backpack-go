package main

import "fmt"

// Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
// Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes
// cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo.
// Dentro de los préstamos que otorga no les cobrará interés a los que posean un sueldo superior a $100.000.
// Es necesario realizar una aplicación que reciba  estas variables y que imprima un mensaje de acuerdo a cada caso.
// Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.

func main() {

	var (
		age       int     = 25
		antiquity int     = 2
		employee  bool    = true
		salary    float64 = 200000.0
	)

	switch {
	case age <= 22:
		fmt.Println("You are under 23")
	case !employee:
		fmt.Println("You are unemployed")
	case antiquity < 1:
		fmt.Println("Your age is less than one year")
	case salary < 100000:
		fmt.Println("Salary is not enough")
	default:
		fmt.Println("You can receive the loan")
	}

}
