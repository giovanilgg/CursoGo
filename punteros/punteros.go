package main

import "fmt"

func main() {
	cadena := "Hola Mundo desde Go"
	fmt.Printf("%p\n", &cadena)
	fmt.Println("Antes de la funcion", cadena)
	modificarValor(&cadena) //pasa una copia y no una referencia de la memoria
	fmt.Println("Despues de la funcion", cadena)

}
func modificarValor(cadena *string) {
	fmt.Printf("%p\n", cadena)
	*cadena = "Hola desde la funcion" //poner asterisco porque estas apuntando a la
	//referncia de la memoria

}
