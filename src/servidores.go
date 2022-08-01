package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	inicio := time.Now()
	servidores := []string{"http://platzi.com", "http://google.com",
		"http://facebook.com", "http://instagram.com"}
	for _, servidor := range servidores {
		revisarServidor(servidor)
	}
	tiempoPaso := time.Since(inicio)
	fmt.Println("tiempo ejecucion %s\n", tiempoPaso)

}

func revisarServidor(servidor string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "No esta disponible =(")
	} else {
		fmt.Println(servidor, "Funciona correctamente")
	}
}
