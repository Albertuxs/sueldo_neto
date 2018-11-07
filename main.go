package main

import (
	c "curso_go/calculo"
	"fmt"
)

var eleccion int

func main() {

	menu :=
		`
	Seleccion tipo de empleo
	[1] Empleado Publico
	[2] Empleado Privado
	`

	fmt.Println("menu", menu)
	fmt.Scanln(&eleccion)


	switch eleccion {
	case 1:
		c.EmpleadoPublico()
	case 2:
		c.EmpleadoPublico()
	default:
		fmt.Println("La opcion es incorrexta")

	}

}

// Aquila clases
// 8 X mes
// 12 X mes
// Pase libre
//
