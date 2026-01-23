package main

import (
	"fmt"
	"practicaGo/models"
	"practicaGo/utils"
)

func main() {

	var Opcion int

	for {
		fmt.Println("\n1 ingresar al sistema")
		fmt.Println("2. Salir")
		fmt.Println("3. Mostrar usuarios")
		fmt.Println("Seleccione una opcion: ")
		fmt.Scanln(&Opcion)

		switch Opcion {
		case 1:

			var p models.Persona

			fmt.Println("ingresa tu nombre: ")
			fmt.Scanln(&p.Name)

			fmt.Println("ingresa tu edad: ")
			fmt.Scanln(&p.Age)

			mensaje := utils.ObtenerCategoria(p.Name, p.Age)
			p.Categoria = mensaje

			utils.GuardarJson(&p)

		case 2:
			fmt.Println("Saliendo del Sistema. Vuelva Prontos")
			return

		case 3:
			utils.ListarUsuarios()

		case 4:
			go utils.IniciarServidor()

		default:
			fmt.Println("opcion invalida")

		}
	}
}
