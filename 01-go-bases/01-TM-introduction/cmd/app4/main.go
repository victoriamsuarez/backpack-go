package main

func main() {

	// Un estudiante de programación intentó realizar declaraciones de variables con sus respectivos tipos en Go, pero tuvo varias dudas mientras lo hacía. A partir de esto, nos brindó su código y pidió la ayuda de un desarrollador experimentado que pueda:
	// Verificar su código y realizar las correcciones necesarias.

	//   var apellido string = "Gomez"	<- Correcta
	//   var edad int = "35"   	 		<- Incorrecta, declara un int y le da valor de un string
	//   boolean := "false";			<- Correcta, puede llegar a prestar la confusión porque se confunde con un booleano, pero en go para declarar se usa la palabra bool. En este caso, es una variable llamada boolean que tiene declarado el string "false"
	//   var sueldo string = 45857.90	<- Incorrecta, se declara una variable tipo string y le da valor de un float
	//   var nombre string = "Julián"	<- Correcta

	// Corregidos

	// var apellido string = "Gomez"
	// var edad int = 35
	// boolean := "false";
	// var sueldo string = "45857.90"
	// var nombre string = "Julián"
}
