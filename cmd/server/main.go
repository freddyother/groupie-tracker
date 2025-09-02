package main

import (
	"groupie-tracker/internal/api"
	"log"
	"net/http"
)

func main() {
	// Rutas para páginas HTML
	http.HandleFunc("/api/artists", api.GetArtists)
	http.HandleFunc("/api/locations", api.GetLocations)
	http.HandleFunc("/api/dates", api.GetDates)
	http.HandleFunc("/api/relation", api.GetRelations)

	// Archivos estáticos
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// Página principal
	http.Handle("/", http.FileServer(http.Dir("web/templates")))

	// Arrancar servidor
	log.Println("🚀 Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
