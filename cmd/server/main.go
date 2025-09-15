package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	api "groupie-tracker/internal/models"
)

const perPage = 3 // show 3 artist per page

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// cargar plantilla
	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		http.Error(w, "Error cargando plantilla", http.StatusInternalServerError)
		return
	}

	// get artist data
	artists := api.GetArtistsData()
	total := len(artists)

	// get page number (default = 1)
	page := 1
	if p := r.URL.Query().Get("page"); p != "" {
		page, _ = strconv.Atoi(p)
		if page < 1 {
			page = 1
		}
	}

	// calculate rangs
	start := (page - 1) * perPage
	end := start + perPage
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}

	// preparar datos para la plantilla
	data := struct {
		Artists  []api.Artist
		Page     int
		HasPrev  bool
		HasNext  bool
		PrevPage int
		NextPage int
	}{
		Artists:  artists[start:end],
		Page:     page,
		HasPrev:  page > 1,
		HasNext:  end < total,
		PrevPage: page - 1,
		NextPage: page + 1,
	}

	// render with data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error ejecutando plantilla", http.StatusInternalServerError)
	}
}

func main() {
	// API JSON
	http.HandleFunc("/api/artists", api.GetArtists)
	http.HandleFunc("/api/locations", api.GetLocations)
	http.HandleFunc("/api/dates", api.GetDates)
	http.HandleFunc("/api/relation", api.GetRelations)

	// Main page
	http.HandleFunc("/", homeHandler)

	// static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	log.Println("🚀 Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
