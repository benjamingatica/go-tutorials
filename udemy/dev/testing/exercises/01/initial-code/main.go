package main

import (
	"fmt"
	"initial-code/perrito"
)

type canino struct {
	nombre string
	edad   int
}

func main() {
	fido := canino{
		nombre: "Fido",
		edad:   perrito.Edad(10),
	}
	fmt.Println(fido)
	fmt.Println(perrito.EdadDos(20))
}
