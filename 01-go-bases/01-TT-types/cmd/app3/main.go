package main

import "fmt"

// Realizar una aplicación que reciba  una variable con el número del mes.
// Según el número, imprimir el mes que corresponda en texto.
// ¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
// Ej: 7, Julio.
// Nota: Validar que el número del mes, sea correcto.

/* Creo que ambas resoluciones están bien aunque cada una tiene sus ventajas. Decidí quedarme con el
mapeo porque hace más legible el código y en este caso no se necesitaba hacer validaciones (si necesitaba
hacer más validaciones usaría el switch ya que el mapeo no me lo permite), en este caso con switch alarga
mucho el código */

func main() {
	fmt.Println(month(14))
}

func month(number int) string {

	// Resolución 1

	// El if valida que los números vayan del 1 al 12, sino manda un mensaje de error
	if number >= 1 && number <= 12 {

		// Inicializo un mapeo donde asocia un int (su clave) con un string
		month := map[int]string{
			1:  "January",
			2:  "February",
			3:  "March",
			4:  "April",
			5:  "May",
			6:  "June",
			7:  "July",
			8:  "August",
			9:  "September",
			10: "October",
			11: "November",
			12: "December",
		}

		// Retorna el string de month asignandole la clave como parámetro
		return month[number]
	} else {
		return "Enter a number from 1 to 12"
	}

	// Resolución 2

	// Inicializa la variable month que devuelve un string
	// var month string

	// Por cada número se le asocia un valor a la variable month. Si falla, tira mensaje de error
	// switch number {
	// case 1:
	// 	month = "January"
	// case 2:
	// 	month = "February"
	// case 3:
	// 	month = "March"
	// case 4:
	// 	month = "April"
	// case 5:
	// 	month = "May"
	// case 6:
	// 	month = "June"
	// case 7:
	// 	month = "July"
	// case 8:
	// 	month = "August"
	// case 9:
	// 	month = "September"
	// case 10:
	// 	month = "October"
	// case 11:
	// 	month = "November"
	// case 12:
	// 	month = "December"
	// default:
	// 	return "Enter a number from 1 to 12"
	// }

	// return month

}
