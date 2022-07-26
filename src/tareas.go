package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t task) marcarCompleta() {
	t.completado = true
}

func (t task) actualizarDescripcion(des string) {
	t.descripcion = des
}

func (t task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t := task{
		nombre:      "Completar mi curso de go",
		descripcion: "Esta semana",
	}
	fmt.Println(t)
	fmt.Printf("%+v\n", t)
	t.marcarCompleta()
	t.actualizarNombre("Finalizar mi curso de go")
	t.actualizarDescripcion("Ahora mismo")
	fmt.Printf("%+v\n", t)
}
