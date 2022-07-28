package main

import "fmt"

type perro struct{}

type pez struct{}

type pajaro struct{}

func (p perro) caminar() string {
	return "Soy un perro y camino"
}

func (p pez) nada() string {
	return "Soy un Pez y estoy nadando"
}

func (p pajaro) vuela() string {
	return "Soy un pajaro y estoy volando"
}

func moverPerro(p perro) {
	fmt.Println(p.caminar())
}

func moverPez(p pez) {
	fmt.Println(p.nada())
}

func moverPajaro(p pajaro) {
	fmt.Println(p.vuela())
}

func main() {

	p := perro{}
	pe := pez{}
	pa := pajaro{}

	moverPerro(p)
	moverPez(pe)
	moverPajaro(pa)

}
