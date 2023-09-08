package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("run finished")
	}()
	file, err := os.Open("./customers.txt")
	if err != nil {
		panic("the indicated file was not found or is damaged")
	}

	defer file.Close()
}
