package main

import "fmt"

// La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla. Para eso tendrán que:
// Crear una aplicación que reciba  una variable con la palabra e imprimir la cantidad de letras que contiene la misma.
// Luego, imprimí cada una de las letras.

func main() {

	// Declaro variable con una palabra
	var word string = "Victoria"

	// len -> funcion nativa que muestra el número de caracteres de una cadena.
	// En una variable guardo la cantidad de caracteres de la variable creada anteriormente
	length := len(word)
	fmt.Println("La cantidad de letras de esta palabra es: ", length)

	// range -> es una forma del ciclo for que itera sobre un slice o un map
	// Uso un for para iterar sobre la palabra declarada en la variable y guardo cada iteración en la variable letter
	for _, letter := range word {
		fmt.Println("Letra: ", string(letter))
	}

}
