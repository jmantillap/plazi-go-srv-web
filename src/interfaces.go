package main

import "fmt"

type animal interface {
	mover() string
}

type perro struct{}

type pez struct{}

type pajaro struct{}

func (p perro) mover() string {
	return "Soy un perro y camino"
}

func (p pez) mover() string {
	return "Soy un Pez y estoy nadando"
}

func (p pajaro) mover() string {
	return "Soy un pajaro y estoy volando"
}

func moverAnimal(a animal) {
	fmt.Println(a.mover())
}

func main() {

	p := perro{}
	pe := pez{}
	pa := pajaro{}

	moverAnimal(p)
	moverAnimal(pe)
	moverAnimal(pa)

}
