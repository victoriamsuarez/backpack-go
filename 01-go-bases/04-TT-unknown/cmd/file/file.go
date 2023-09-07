package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		fmt.Println("Initializing...")
		fmt.Println("run finished")
	}()
	file, err := os.Open("./customers.txt")
	if err != nil {
		panic("the indicated file was not found or is damaged")
	}

	defer file.Close()
}
