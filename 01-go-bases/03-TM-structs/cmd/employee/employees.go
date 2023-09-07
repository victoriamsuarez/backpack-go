package main

import "fmt"

// Crear una estructura Person con los campos ID, Name, DateOfBirth.
type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

// Crear una estructura Employee con los campos: ID, Position y una composicion con la estructura Person.
type Employee struct {
	ID         int
	Position   string
	DatePerson Person
}

func main() {

	// Instanciar en la función main() tanto una Person como un Employee cargando sus respectivos campos y por último ejecutar el método PrintEmployee().
	pers := Person{
		ID:          100,
		Name:        "Victoria",
		DateOfBirth: "29-05-1998",
	}

	empl := Employee{
		ID:         1,
		Position:   "software developer",
		DatePerson: pers,
	}

	empl.PrintEmployee()

}

// Realizar el método a la estructura Employe que se llame PrintEmployee(), lo que hará es realizar la impresión de los campos de un empleado.
func (e Employee) PrintEmployee() {
	fmt.Println(e)
}
