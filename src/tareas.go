package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func main() {
	t := task{
		nombre:      "Completar mi curso de go",
		descripcion: "Esta semana",
	}

	fmt.Println(t)
	fmt.Printf("%+v\n", t)
}
