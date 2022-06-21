package main

import (
	"01-codigo-inicial/cita"
	"01-codigo-inicial/palabra"
	"fmt"
)

func main() {
	fmt.Println(palabra.Conteo(cita.Cuando))

	for k, v := range palabra.ConteoUso(cita.Cuando) {
		fmt.Println(v, k)
	}
}
