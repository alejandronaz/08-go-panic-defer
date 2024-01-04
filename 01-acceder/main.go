package main

import (
	"fmt"
	"os"
)

// file, err := os.Open("archivo.txt")
// bytes, err := io.ReadAll(file)
// data := string(bytes)

func main() {

	_, err := os.Open("archivo.txt")

	defer func() {
		fmt.Println("Ejecucion finalizada")
	}()

	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

}
