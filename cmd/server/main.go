package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("🚀 Servidor iniciado en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error iniciando servidor: ", err)
	}
}
