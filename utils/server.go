package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"practicaGo/models"
)

func ListarUsuariosWeb(w http.ResponseWriter, r *http.Request) {
	archivo := "usuarios.json"
	var lista []models.Persona

	datos, err := os.ReadFile(archivo)
	if err != nil {
		http.Error(w, "no se pudo leer el archivo", http.StatusInternalServerError)
		return
	}

	json.Unmarshal(datos, &lista)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lista)
}

func IniciarServidor() {
	http.HandleFunc("/usuarios", ListarUsuariosWeb)

	fmt.Println("server run")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error al iniciar servidor", err)
	}
}
