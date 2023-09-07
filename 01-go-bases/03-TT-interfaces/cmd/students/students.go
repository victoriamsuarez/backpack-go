package main

import "fmt"

// Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:
// Nombre: [Nombre del alumno]
// Apellido: [Apellido del alumno]
// DNI: [DNI del alumno]
// Fecha: [Fecha ingreso alumno]
// Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
// Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle

type Student struct {
	FirtName string
	LastName string
	DNI      string
	Date     string
}

func main() {

	// Creo un nuevo estudiudiante. Un puntero de Student
	student := &Student{}

	// Llamo método para inicializarlo
	student.createNewStudent("victoria", "suarez", "12345678", "22-08-2023")

	// Muestro detalle en consola
	student.detail()

}

// var Students = []Student{}

func (s *Student) detail() {
	fmt.Printf("FirtName: %s \nLastName: %s\nDNI: %s \nDate: %s \n", s.FirtName, s.LastName, s.DNI, s.Date)
}

func (s *Student) createNewStudent(name, lastname, dni, date string) {
	s.FirtName = name
	s.LastName = lastname
	s.DNI = dni
	s.Date = date
}
