package main

import "fmt"

/* Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. No se pueden introducir notas negativas. */

func main() {
	student := []int{9, 8, 7, 9, 9, 8, 7, 10}
	fmt.Println(calculateAverage(student))
}

func calculateAverage(studentRatings []int) float64 {
	// Inicializo para guardar la suma de calificaciones
	sum := 0

	// range indica que va a iterar sobre un slice (en este caso studentRatings) o map en un bucle for
	// Este for recorre el range de studentRatings. Rating toma cada elemento de studentRatings, los suma entre sí y los guarda en la variable sum
	for _, ratings := range studentRatings {
		sum += ratings
	}

	// len da la longitud del studentRatings y si es mayor a cero, retorna la variable sum dividido la longitud de studentRatings para sacar el promedio
	if len(studentRatings) > 0 {
		return float64(sum) / float64(len(studentRatings))
	} else {
		return 0.0
	}
}
