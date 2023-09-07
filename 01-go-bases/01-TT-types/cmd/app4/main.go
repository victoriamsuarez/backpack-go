package main

import "fmt"

func main() {

	/* Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
	   Según el siguiente mapa, ayudá a imprimir la edad de Benjamin.

	     var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	   Por otro lado también es necesario:
	   Saber cuántos de sus empleados son mayores de 21 años.
	   Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	   Eliminar a Pedro del mapa.
	*/

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	var count int

	// Cuantos empleados tienen más de 21 años
	for _, age := range employees {
		if age > 21 {
			count++
		}
	}

	fmt.Println("There are", count, "employees over 21 years old")

	// Agrega un nuevo employee
	employees["Federico"] = 25

	fmt.Println("New employee Federico")
	fmt.Println(employees)

	// Eliminar un employee
	delete(employees, "Pedro")

	fmt.Println("Delete employee Pedro")
	fmt.Println(employees)
}
