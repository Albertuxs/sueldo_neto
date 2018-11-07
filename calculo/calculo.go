package calculo

import "fmt"

var sueldo int
var gremio int
var salario int
var neto int



func EmpleadoPublico() {
	prome := 17
	// Pregunta y guarda los datos
	fmt.Println("Ingrese su sueldo bruto: ")
	fmt.Scanln(&sueldo)
	fmt.Println("Tiene descuento gremial, cual es su porcentaje:")
	fmt.Scanln(&gremio)

	// El if es si tiene gremio agrega el descuento adicional al sueldo
	if gremio > 0 { // Resultado con gremio
		discount := &gremio + &prome
		gremial := ((&sueldo * &discount) / 100)
		fmt.Println("Su sueldo neto con descuento gremial es: ", gremial)


	} else { // Resultado sin gremio
		&neto = ((sueldo * &prome) / 100)
		&salario = &prome + sueldo

	}
	fmt.Println("Su sueldo es: ", salario)
}

func EmpleadoPrivado() {
	fmt.Println("Ingrese su sueldo: ")
	fmt.Scanln(&sueldo)
	promedio := ((sueldo * 17) / 100)
	salario := promedio + sueldo
	fmt.Println("Su sueldo es: ", salario)
}
