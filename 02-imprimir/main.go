package main

import (
	"fmt"
	"io"
	"os"
)

// file, err := os.Open("archivo.txt")
// bytes, err := io.ReadAll(file)
// data := string(bytes)

func main() {

	file, err := os.Open("customers.txt")

	defer func() {
		fmt.Println("Ejecucion finalizada")
	}()

	defer file.Close()

	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic("Cant read the file")
	}
	fmt.Println(string(bytes))

}
