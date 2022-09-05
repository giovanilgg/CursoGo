package main

import (
	"fmt"
	"os"
)

func main() {
	//open devulve dos valores

	defer func() {
		//la funcion recover controla los panic en caso que se desencadene un evento
		if error := recover(); error != nil {
			fmt.Println("al parecer el programa no finalizo de manera correcta")

		}
	}()

	if file, error := os.Open("hola.txt"); error != nil {
		panic("no se pudo leer el archivo") //mensaje de errores

	} else {
		defer func() { //defer se ejecuta despues del main
			file.Close() //cerra el archivo

		}()
		contenido := make([]byte, 254)                //el file trae un archivo binario de tipo byte
		long, _ := file.Read(contenido)               //read necesita donde guardar y un error
		contenidoArchivo := string(contenido[0:long]) //parsea a string el binario
		fmt.Println(contenidoArchivo)

	}

}
