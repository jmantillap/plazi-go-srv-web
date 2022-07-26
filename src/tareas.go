package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarLista(tarea *task) {
	t.tasks = append(t.tasks, tarea)
}

func (t *taskList) eliminarLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(des string) {
	t.descripcion = des
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t1 := &task{
		nombre:      "Completar mi curso de go",
		descripcion: "Esta semana",
	}
	t2 := &task{
		nombre:      "Completar mi curso de pyton",
		descripcion: "Esta semana",
	}
	t3 := &task{
		nombre:      "Completar mi curso de nodejs",
		descripcion: "Esta semana",
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}
	fmt.Println(lista.tasks[0])
	lista.agregarLista(t3)
	fmt.Println(len(lista.tasks))
	lista.eliminarLista(1)
	fmt.Println(len(lista.tasks))

}
