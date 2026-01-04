package main

import (
	"fmt"
	"os"
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

func guardarRegistro(texto string) {
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error al abrir archivo:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(texto + "\n")

	if err != nil {
		fmt.Println("error al escribir:", err)
	}
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

			guardarRegistro("\n Nombre:" + Name + "Resultado: " + mensaje)

		case 2:
			fmt.Println("Saliendo del Sistema. Vuelva Prontos")
			return
		default:
			fmt.Println("opcion invalida")

		}
	}
}
