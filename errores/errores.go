package main

import (
	"errors"
	"fmt"
)

//variables globales fuera de la funcion
// defer gio()  la funcion gio se ejecuta despues del main
func main() {
	result, error := division(4, 1)
	if error == nil {
		fmt.Println("el resultado es igual", result)
	} else {
		fmt.Println("el error es ", error)

	}

}
func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir entre 0")

	} else {

		return dividendo / divisor, nil
	}
}
