package main

import "fmt"

// Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func main() {

	//Ejecutar al menos una vez cada método y función definido desde main().
	p1 := Product{
		ID:          1,
		Name:        "cereal bar",
		Price:       160.0,
		Description: "it is a healthy snack",
		Category:    "snack",
	}

	p1.GetAll()
	p1.Save()
	fmt.Println("Added product")
	p1.GetAll()

	fmt.Println("Search by ID")
	searchById := getById(1)
	fmt.Println(searchById)

}

// Tener un slice global de Product llamado Products instanciado con valores.
var Products = []Product{
	{ID: 2, Name: "soda", Price: 200.0, Description: "orange flavor drink", Category: "drink"},
	{ID: 4, Name: "juice", Price: 170.0, Description: "multifruit flavor drink", Category: "drink"},
}

// El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método.
func (p *Product) Save() {
	Products = append(Products, *p)
}

// El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
func (p Product) GetAll() {
	for _, prod := range Products {
		fmt.Println(prod)
	}
}

// Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
func getById(id int) *Product {
	for _, p := range Products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}
