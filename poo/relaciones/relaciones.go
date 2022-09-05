package main

import "fmt"

type Usuario struct {
	nombre string
	email  string
	activo bool
}

//relacion 1:1
type Estudiante struct {
	user   Usuario
	codigo int
}

//relacion 1:m

type Curso struct {
	titulo string
	videos []Video //videos arreglo de tipo videos
}
type Video struct {
	titulo string
	curso  Curso
}

func main() {
	us1 := Usuario{
		nombre: "gio",
		email:  "giovanizzxx@gmail.com",
		activo: true,
	}
	us2 := Usuario{
		nombre: "otro",
		email:  "otro@gmail.com",
		activo: true,
	}
	//relacion 1:1
	estudianteGio := Estudiante{
		user:   us1,
		codigo: 418114582,
	}
	//.----------------------------------
	video1 := Video{titulo: "Introduccio a la programacion0"}
	video2 := Video{titulo: "Programacion avanazada"}

	curso := Curso{
		titulo: "cursoGiovas",
		videos: []Video{video1, video2},
	}
	video1.curso = curso
	video2.curso = curso

	fmt.Println(curso)
	fmt.Println(us2)
	fmt.Println(estudianteGio.user.nombre)
}
