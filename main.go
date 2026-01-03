package main

import (
	"fmt"
)

func obtenerCategoria(Name string, Age int) string {
	if Name == "admin" || Name == "Admin" {
		return "acceso total, Administrador"
	}
	if Age < 18 {
		return "acceso denegado, eres menor de edad"
	} else if Age >= 18 && Age <= 30 {
		return "Bienvenido junior"
	}
	return "Bienvenido Senior"
}

func main() {

	var Name string
	var Age int
	var Opcion int

	for {
		fmt.Println("\n1 ingresar al sistema")
		fmt.Println("2. Salir")
		fmt.Println("Seleccione una opcion: ")
		fmt.Scanln(&Opcion)

		switch Opcion {
		case 1:
			fmt.Println("ingresa tu nombre: ")
			fmt.Scanln(&Name)

			fmt.Println("ingresa tu edad: ")
			fmt.Scanln(&Age)

			mensaje := obtenerCategoria(Name, Age)

			fmt.Println(mensaje)

		case 2:
			fmt.Println("Saliendo del Sistema. Vuelva Prontos")
			return
		default:
			fmt.Println("opcion invalida")

		}
	}
}
