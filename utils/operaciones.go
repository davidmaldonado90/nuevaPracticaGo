package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"practicaGo/models"
	"strings"
)

func ObtenerCategoria(Name string, Age int) string {
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

func ListarUsuarios() {
	archivo := "usuarios.json"
	var lista []models.Persona

	datos, err := os.ReadFile(archivo)
	if err != nil {
		fmt.Println("no hay usuarios registrados")
		return
	}

	json.Unmarshal(datos, &lista)

	fmt.Println("\n--- LISTA DE USUARIOS ---")

	for i, u := range lista {
		linea := fmt.Sprintf("%d. Nombre: %-10s | Edad: %d | Categoria: %s", i+1, u.Name, u.Age, u.Categoria)

		fmt.Println(linea)
	}
	fmt.Println("----------------------------------")
}

// func guardarRegistro(texto string) {
// 	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		fmt.Println("error al abrir archivo:", err)
// 		return
// 	}
// 	defer f.Close()

// 	_, err = f.WriteString(texto + "\n")

// 	if err != nil {
// 		fmt.Println("error al escribir:", err)
// 	}
// }

func GuardarJson(p *models.Persona) {
	archivo := "usuarios.json"
	var lista []models.Persona

	datos, err := os.ReadFile(archivo)
	if err == nil {
		json.Unmarshal(datos, &lista)
	}

	for _, usuario := range lista {
		if strings.EqualFold(usuario.Name, p.Name) {
			fmt.Println("el usuario: ", p.Name, "ya existe")
			return
		}
	}

	lista = append(lista, *p)

	datosJson, _ := json.MarshalIndent(lista, "", "  ")

	err = os.WriteFile(archivo, datosJson, 0644)
	if err != nil {
		fmt.Println("error al guardar", err)
	} else {
		fmt.Println("usuario registrado con exito")
	}

}
