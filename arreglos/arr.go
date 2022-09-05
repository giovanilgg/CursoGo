package main

import (
	"fmt"
)

func main() {
	var numeros [5]int
	fmt.Println(numeros)

	//arreglo de tipo string
	nombres := [5]string{"Gio", "Gio2", "Gio3"}
	for i := 0; i < len(numeros); i++ {
		fmt.Println(nombres[0], i)
	}
	//arreglo  sin longitud fija
	colores := [...]string{"rojo", "verde", "amarillo"}

	fmt.Println(len(nombres), len(colores))
	//colocar indices definidos

	monedas := [...]string{
		0:  "pesos",
		3:  "dolares",
		20: "euros",
	}
	for i := 0; i < len(monedas); i++ {
		fmt.Println(len(monedas))
	}
	//creando un subarreglo   indico desde la posicion 1 a la una antes de 5 [:3] final y principio ya no
	//se pone 0 puede ser :4 inicio  4: final
	moneditas := monedas[1:5]
	fmt.Println(moneditas)

	//slicen  es como un arreglo al cual le podemos agregar elementos
	slice := []int{2, 3, 4}
	fmt.Println(slice, len(slice))
	slice = append(slice, 4)
	slice = append(slice, 6)
	subSlice := slice[:2]
	slice[0] = 299292
	//el subslice afecta al padre arreglo por eso subslice toma en cuenta el 2299293
	fmt.Println(slice, len(slice))
	fmt.Println(subSlice, len(subSlice))
	//punteros,longitud y capacidad

	meses := []string{"enero", "febrero", "marzo", "abril"}
	fmt.Printf("Len:%v ,Cap:%v , %p \n", len(meses), cap(meses), meses)
	meses = append(meses, "noviembre")
	fmt.Printf("Len:%v ,Cap:%v , %p \n", len(meses), cap(meses), meses)

	//funcion make para contruir un sliecen ,es como el new
	prendas := make([]int, 3, 3) //primer argumento de que tipo
	//segundo la longitud y tercero la capacidad
	prendas[0] = 100
	prendas[1] = 200
	prendas[2] = 300
	prendas = append(prendas, 400)
	fmt.Println(prendas, len(prendas), cap(prendas))

}
