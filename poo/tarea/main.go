package main

import "fmt"

type ListaTask struct {
	task []*Task //se crea un slicen de tareas de tipo clase task

}

func (tl *ListaTask) agregarTarea(t *Task) {
	tl.task = append(tl.task, t) //almacenamos en mi slicen lo que se pase en la lista de tareas

}

func (tl *ListaTask) removerTarea(index int) {
	tl.task = append(tl.task[:index], tl.task[index+1:]...)
	//excluye el indice actual y retorna una nueva lista

}

type Task struct {
	name      string
	desc      string
	completed bool
}

func (t *Task) imprimir() {

	fmt.Printf("Nombre:%s\n Descripcion:%s\n Completada:%t \n ", t.name, t.desc, t.completed)

}
func (t *Task) Comp() {

	t.completed = true
}
func main() {
	tarea1 := Task{

		name:      "Curso de Go",
		desc:      "Completar curso de Go",
		completed: false,
	}
	tarea2 := Task{

		name:      "Curso de React",
		desc:      "Completar curso de Go",
		completed: false,
	}
	tarea3 := Task{

		name:      "Curso de Node",
		desc:      "Completar curso de Go",
		completed: false,
	}

	listaDeTareas := ListaTask{}
	listaDeTareas.agregarTarea(&tarea1) //con el & no manda copia si no la referencia
	listaDeTareas.agregarTarea(&tarea2)
	listaDeTareas.agregarTarea(&tarea3)

	for _, valor := range listaDeTareas.task {

		//fmt.Println(listaDeTareas.task[i])
		fmt.Println(valor)

	}
	listaDeTareas.removerTarea(0)
	for _, valor := range listaDeTareas.task {

		//fmt.Println(listaDeTareas.task[i])
		fmt.Println(valor)

	}

}
