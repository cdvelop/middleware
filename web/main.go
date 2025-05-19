package main

import (
	"fmt"
	"log"
	"net/http"
)

// servidor web
func main() {

	fs := http.FileServer(http.Dir("./public"))

	http.Handle("/", fs)

	http.HandleFunc("POST /login", Login())

	fmt.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Parsear formulario multipart (importante para FormData)
		maxMemory := int64(10 * 1024 * 1024) // 10MB máximo de memoria para archivos
		err := r.ParseMultipartForm(maxMemory)
		if err != nil {
			http.Error(w, "Error al procesar el formulario: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Extraer datos del formulario
		username := r.FormValue("username")
		password := r.FormValue("password")

		fmt.Fprintf(w, "Usuario: %s, Contraseña: %s\n", username, password)

	}
}
