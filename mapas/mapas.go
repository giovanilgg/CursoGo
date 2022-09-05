package main

import (
	"fmt"
)

func main() {

	//mapas
	dias := make(map[int]string) //par llave valor make contruye un mapa
	//donde tiene la clave entera y el valor es string
	//make solo inicializa slice,mapa o channel
	//a los mapas no se les puede imprimir la capacidad
	//se asignan de nuevo por la llave

	dias[1] = "Lunes"
	dias[2] = "Martes"
	dias[3] = "Miercoles"
	dias[4] = "Jueves"
	dias[5] = "Viernes"
	fmt.Println(dias)
	delete(dias, 1) //sliminar una parte del mapa
	fmt.Println(dias)

	//nuevo mapa
	estudiantes := make(map[string][]int)

	estudiantes["Giovani"] = []int{25, 36}
	fmt.Println(estudiantes["Giovani"])    //acceder a un elemento de un map
	fmt.Println(estudiantes["Giovani"][1]) //acceder a un elemento ,especifico de mi valor
}
