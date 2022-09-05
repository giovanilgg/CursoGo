package main

import (
	"fmt"
	"net/http"
	"time"
)

//revisar si el servidor funciona

func revisarServidor(servidor string, canal chan string) {
	//utilizar paqute http
	_, error := http.Get(servidor)
	if error != nil {

		//fmt.Println("servidor no siponible")
		canal <- servidor + "servidor no siponible"
	} else {
		//fmt.Println("esta funcionando el servidor ")
		canal <- servidor + "esta funcionando"

	}
}

func main() {
	//creando canal   con el el make
	canal := make(chan string)
	//aplicacion sin concurrencia
	inicio := time.Now() //devuleve el tiempp actual
	servidores := []string{
		"https://oregoom.com/",
		"https://www.udemy.com/",
		"https://www.youtube.com/",
		"https://www.facebook.com/",
		"https://www.google.com/",
		"http://localhost:4200/indicios",
	}
	for _, valor := range servidores {

		go revisarServidor(valor, canal) //ejecuta la intruccion en multiples hilos

	}
	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal) //obtiene lo que esta pasando en ese canal

	}
	tiempoTrancurrido := time.Since(inicio)
	fmt.Println(tiempoTrancurrido)

}
