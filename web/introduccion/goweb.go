package main

import (
	"fmt"
	"log"
	"net/http"
)

//handler
func Hola(res http.ResponseWriter, req *http.Request) {
	fmt.Println("el metodo es", req.Method)
	fmt.Fprintln(res, "Hola mundo")
}
func PaginaNE(res http.ResponseWriter, req *http.Request) {
	http.NotFound(res, req) //status 404

}
func Error(res http.ResponseWriter, req *http.Request) {
	//http.Error(res, "La pagina no funciona", 404) //status 404
	http.Error(res, "La pagina no funciona", http.StatusPaymentRequired)
}
func Saludar(res http.ResponseWriter, req *http.Request) {
	//http.Error(res, "La pagina no funciona", 404) //status 404
	fmt.Println(req.URL)
	fmt.Println(req.URL.RawQuery) //envia los datos de la url
	fmt.Println(req.URL.Query())  //cacha lod datos de la url y lo converte en un map
	name := req.URL.Query().Get("name")
	fmt.Println(name)
}

func main() {
	//creand el mux

	mux := http.NewServeMux()

	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/page", PaginaNE)
	mux.HandleFunc("/error", Error)
	mux.HandleFunc("/saludar", Saludar)

	//creando la ruta
	/*http.HandleFunc("/", Hola)
	http.HandleFunc("/page", PaginaNE)
	http.HandleFunc("/error", Error)
	http.HandleFunc("/saludar", Saludar)*/

	//creando el server

	server := &http.Server{

		Addr:    "localhost:8000",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
	/*fmt.Println("el servidor esta corriendo en el puerto 8000")
	fmt.Println("Run server:http://localhost:8000/ ")
	log.Fatal(http.ListenAndServe("localhost:8000", mux))*/

}
