package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

//estructura de usuari
type Usuarios struct {
	Nombre string
	Correo string
	Edad   int
	Activo bool
	Admin  bool
	Cursos []Curso
}
type Curso struct {
	Nombre string
}

//... los tres puntos significa que recibe n cosas
func Saludar(nombre string) string {

	return "Hola desde una funcion" + nombre
}

func (u *Usuarios) Contructor(nombre string, correo string, edad int, activo bool, admin bool, cursos []Curso) {
	u.Nombre = nombre
	u.Correo = correo
	u.Edad = edad
	u.Activo = activo
	u.Admin = admin
	u.Cursos = cursos

}

//usando Execute template para usarlo globalmete
var funciones = template.FuncMap{"saludo": Saludar}
var templates = template.Must(template.New("").Funcs(funciones).ParseGlob("templates/**/*.html"))

func renderizarTemplate(res http.ResponseWriter, plantilla string, datos interface{}) {
	error := templates.ExecuteTemplate(res, plantilla, datos)
	if error != nil {

		http.Error(res, "no es posible teronar el template", http.StatusInternalServerError)
	}
}
func index(res http.ResponseWriter, req *http.Request) {
	persona := Usuarios{}
	c1 := Curso{"java"}
	c2 := Curso{"javascript"}
	c3 := Curso{"python"}
	c4 := Curso{"ruby"}
	c5 := Curso{"Perl"}
	cursos := []Curso{c1, c2, c3, c4, c5}
	persona.Contructor("Giovani", "giovanilopez82@aragon.unam.mx", 25, true, true, cursos)

	//con must maneja en automatico los errores a diferencia de parsefiles que te retorna
	//el error y lo tienes que hacer manual
	/*---------------------------------------------------bloque anterior
	Template, error := template.New("index.html").Funcs(funciones).ParseFiles("index.html", "molde.html") //ruta de nuestros archivos
	if error != nil {
		panic("No se encuentra la plantilla a renderizar ")
	} else {
		Template.Execute(res, persona) //recibe recibe la plantilla y una estructura
	}*/
	//bloque nuevo -------------------------------------7

	renderizarTemplate(res, "index.html", persona)

}
func nosotros(res http.ResponseWriter, req *http.Request) {
	renderizarTemplate(res, "nosotros.html", nil)

}

func main() {

	//Archivos estaticos
	statucFile := http.FileServer(http.Dir("static"))
	//creando un usuario

	//creando las rutas
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/nosotros", nosotros)
	//mux de static files
	mux.Handle("/static/", http.StripPrefix("/static/", statucFile))

	//creando el server
	fmt.Println("corrien el server en el puerto 8000")
	server := &http.Server{

		Addr:    "localhost:8000",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
