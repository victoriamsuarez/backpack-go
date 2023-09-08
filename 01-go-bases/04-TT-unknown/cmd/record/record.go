package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type Client struct {
	Docket      int    `json:"docket:"`
	Name        string `json:"name:"`
	ID          int    `json:"id:"`
	PhoneNumber string `json:"phonenumber:"`
	Residence   string `json:"residence:"`
}

var clientsGroup []Client = []Client{

	{Docket: 1, Name: "Robert", ID: 12345678, PhoneNumber: "457683923", Residence: "resience 123"},
	{Docket: 2, Name: "Elisabeth", ID: 27683453, PhoneNumber: "237856226", Residence: "residence 456"},
	{Docket: 3, Name: "Federico", ID: 0, PhoneNumber: "928346725", Residence: "residence789"},
}

func printClient() {

	for _, c := range clientsGroup {
		jsonC, err := json.MarshalIndent(c, "", " ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonC))
	}

}

func registerCustomer(c Client) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println("several errors were detected at runtime")
			fmt.Println("run finished")
		}
	}()

	_, err := existingCustomer(c)
	if err != nil {
		panic(err)
	} else {
		_, err := c.validateData()
		if err != nil {
			panic(err)
		} else {
			clientsGroup = append(clientsGroup, c)

			file, err := os.Open("customers.txt")
			if err != nil {
				panic("error al abrir el archivo")
			}

			data, err := json.MarshalIndent(clientsGroup, "", " ")
			if err != nil {
				panic("error format json")
			}

			err = json.Unmarshal(data, &clientsGroup)
			if err != nil {
				panic("error al decoficar")
			}

			err = os.WriteFile("customers.txt", data, 0644)
			if err != nil {
				panic("the indicated file was not found or is damaged")
			}
			defer file.Close()
		}

	}

}

func existingCustomer(client Client) (bool, error) {

	for _, existing := range clientsGroup {
		if existing.Docket == client.Docket {
			return false, errors.New("client already exists")
		}
	}
	return true, nil

}

func (client Client) validateData() (bool, error) {
	if client.Docket == 0 || client.Name == "" || client.ID == 0 || client.PhoneNumber == "" || client.Residence == "" {
		return false, errors.New("error: cannot have null fields")
	}
	return true, nil
}
func main() {

	fmt.Println("Lista de clientes principal:")

	printClient()

	fmt.Println("Registo de cliente")

	c1 := Client{4, "Rufina", 56708676, "78647986", "Home 765"}

	registerCustomer(c1)

	fmt.Println("Lista de usuarios completa:")

	printClient()

	fmt.Println("Hasta acá registro un cliente y lo imprimo para verificar")

	fmt.Println("Registo de cliente. Solo archivo, no muestra por consola")

	c4 := Client{8, "Emanuel", 89435693, "29385684235", "Home 901"}

	registerCustomer(c4)

	// fmt.Println("Prueba cliente ya registrado")

	// registerCustomer(c1)

	// fmt.Println("Prueba de clientes con campos nulos")

	// c2 := Client{5, "Gilda", 8744566, "", ""}

	// registerCustomer(c2)

	// fmt.Println("Prueba de clientes con legajo repetido")

	// c3 := Client{4, "Pedro", 56708676, "78647986", "Home 765"}

	// registerCustomer(c3)
}
