package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

// Ejemplo de estructura (normalmente vendrá de la API externa)
type Artist struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Members []string `json:"members"`
}

// Simulamos datos para la demo
var artists = []Artist{
	{ID: 1, Name: "Pink Floyd", Members: []string{"David Gilmour", "Roger Waters"}},
	{ID: 2, Name: "The Beatles", Members: []string{"John Lennon", "Paul McCartney", "George Harrison", "Ringo Starr"}},
}

// Handler: Home (renderiza HTML)
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		http.Error(w, "Error cargando template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, artists) // Pasamos datos a la vista
	if err != nil {
		http.Error(w, "Error ejecutando template", http.StatusInternalServerError)
	}
}

// Handler: API JSON con lista de artistas
func apiArtistsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(artists)
	if err != nil {
		http.Error(w, "Error codificando JSON", http.StatusInternalServerError)
	}
}

func main() {
	// Rutas para páginas HTML
	http.HandleFunc("/", homeHandler)

	// Rutas para API JSON
	http.HandleFunc("/api/artists", apiArtistsHandler)

	// Servir archivos estáticos (CSS, JS, imágenes)
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Arrancar servidor
	log.Println("🚀 Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
