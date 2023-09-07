package main

import (
	"errors"
	"fmt"
	"math"
)

/* Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los/as estudiantes de un curso. Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.
Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior. */

func main() {

	ratings := []int{8, 9, 7, 5, 4, 8, 9, 10}

	const (
		minimum = "minimum"
		average = "average"
		maximum = "maximum"
	)

	minValue, err := operation(minimum, ratings)
	averageValue, err := operation(average, ratings)
	maxValue, err := operation(maximum, ratings)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Minimum rating:", minValue)
	fmt.Println("Average ratings:", averageValue)
	fmt.Println("Maximum rating:", maxValue)
}

// Creo una función que reciba como parámetro un string de lo que se solicita y un slice de números
func operation(function string, numbers []int) (float64, error) {

	// Uso un switch para evaluar el string que viene desde el parámetro y le asigno a cada función como parámetro un slice de números. En caso de que no sea un pámetro esperado, muestra un mensaje de error y terorna un math.NaN() ya que no existe una gunción que realice ese tipo de cáculo
	switch function {
	case "minimum":
		return float64(minFunc(numbers)), nil
	case "average":
		return averageFunc(numbers), nil
	case "maximum":
		return float64(maxFunc(numbers)), nil
	default:
		return math.NaN(), errors.New("invalid function. try with minimum, maximum or average")
	}

}

func minFunc(numbers []int) int {
	minimum := numbers[0]

	for _, number := range numbers {
		if number < minimum {
			minimum = number
		}
	}
	return minimum

}

func averageFunc(numbers []int) float64 {
	total := 0

	for _, ratings := range numbers {
		total += ratings
	}

	if len(numbers) > 0 {
		return float64(total) / float64(len(numbers))
	} else {
		return 0
	}
}

func maxFunc(numbers []int) int {
	maximum := numbers[0]

	for _, number := range numbers {
		if number > maximum {
			maximum = number
		}
	}
	return maximum
}
