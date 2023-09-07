package main

import (
	"fmt"

	product "github.com/victoriamsuarez/backpack-go/cmd/internal"
)

// Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
// La empresa tiene 3 tipos de productos: Pequeño, Mediano y Grande. (Se espera que sean muchos más)

// Y los costos adicionales son:
// Pequeño: solo tiene el costo del producto
// Mediano: el precio del producto + un 3% de mantenerlo en la tienda
// Grande: el precio del producto + un 6% de mantenerlo en la tienda, y adicional a eso $2500 de costo de envío.

// El porcentaje de mantenerlo en la tienda es en base al precio del producto.
// El costo de mantener el producto en stock en la tienda es un porcentaje del precio del producto.

// Se requiere una función factory que reciba el tipo de producto y el precio y retorne una interfaz Producto que tenga el método Precio.

// Se debe poder ejecutar el método Precio y que el método me devuelva el precio total en base al costo del producto y los adicionales en caso que los tenga

func main() {

	configS := &product.ConfigSmall{
		Price: 1000.0,
	}

	configM := &product.ConfigMedium{
		Price: 1000.0,
	}

	configL := &product.ConfigLarge{
		Price: 1000.0,
	}

	prod1 := newProduct(ProdSmall, configS)
	prod2 := newProduct(ProdMedium, configM)
	prod3 := newProduct(ProdLarge, configL)

	fmt.Println(prod1.Price())
	fmt.Println(prod2.Price())
	fmt.Println(prod3.Price())

}

const (
	ProdSmall  = "small"
	ProdMedium = "medium"
	ProdLarge  = "largue"
)

func newProduct(prodType string, config any) product.Product {
	switch prodType {
	case ProdSmall:
		p := config.(*product.ConfigSmall)
		return product.NewSmall(p)
	case ProdMedium:
		p := config.(*product.ConfigMedium)
		return product.NewMedium(p)
	case ProdLarge:
		p := config.(*product.ConfigLarge)
		return product.NewLargue(p)
	default:
		return nil
	}
}
