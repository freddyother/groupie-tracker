package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	api "groupie-tracker/internal/models"
)

const perPage = 3 // show 3 artist per page

func renderFile(w http.ResponseWriter, file string, data any) {
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		log.Println("template error:", err)
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Println("exec error:", err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

// / or /artists
func artistsHandler(w http.ResponseWriter, r *http.Request) {
	artists := api.GetArtistsData()
	total := len(artists)

	page := 1
	if p := r.URL.Query().Get("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}
	start := (page - 1) * perPage
	end := start + perPage
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}

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

	renderFile(w, "web/templates/index.html", data)
}

// helpers
func artistNameByID(id int) string {
	for _, a := range api.GetArtistsData() {
		if a.ID == id {
			return a.Name
		}
	}
	return "Unknown"
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func locationsHandler(w http.ResponseWriter, r *http.Request) {
	// view model
	type LocVM struct {
		Name      string
		Locations []string
	}

	// construir lista con nombre de artista
	var items []LocVM
	for _, l := range api.GetLocationsData() { // crea esta función como GetArtistsData() o usa locations directamente si la tienes pública
		items = append(items, LocVM{
			Name:      artistNameByID(l.ID),
			Locations: l.Locations,
		})
	}

	// paginación
	page := 1
	if p := r.URL.Query().Get("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}
	total := len(items)
	start := (page - 1) * perPage
	if start > total {
		start = total
	}
	end := min(start+perPage, total)

	data := struct {
		Items    []LocVM
		Page     int
		HasPrev  bool
		HasNext  bool
		PrevPage int
		NextPage int
	}{items[start:end], page, page > 1, end < total, page - 1, page + 1}

	renderFile(w, "web/templates/locations.html", data)
}

func datesHandler(w http.ResponseWriter, r *http.Request) {
	type DateVM struct {
		Name  string
		Dates []string
	}

	var items []DateVM
	for _, d := range api.GetDatesData() { // crea GetDatesData() análogo o usa tu slice si es público
		items = append(items, DateVM{
			Name:  artistNameByID(d.ID),
			Dates: d.Dates,
		})
	}

	page := 1
	if p := r.URL.Query().Get("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}
	total := len(items)
	start := (page - 1) * perPage
	if start > total {
		start = total
	}
	end := min(start+perPage, total)

	data := struct {
		Items    []DateVM
		Page     int
		HasPrev  bool
		HasNext  bool
		PrevPage int
		NextPage int
	}{items[start:end], page, page > 1, end < total, page - 1, page + 1}

	renderFile(w, "web/templates/dates.html", data)
}

func relationHandler(w http.ResponseWriter, r *http.Request) {
	type RelPair struct {
		Place string
		When  []string
	}
	type RelVM struct {
		Name       string
		Pairs      []RelPair
		Cities     int
		TotalDates int
	}

	var items []RelVM
	for _, rel := range api.GetRelationsData() { // crea GetRelationsData() o expón tu slice
		var pairs []RelPair
		totalDates := 0
		for place, when := range rel.Relations {
			pairs = append(pairs, RelPair{Place: place, When: when})
			totalDates += len(when)
		}
		items = append(items, RelVM{
			Name:       artistNameByID(rel.ID),
			Pairs:      pairs,
			Cities:     len(rel.Relations),
			TotalDates: totalDates,
		})
	}

	page := 1
	if p := r.URL.Query().Get("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}
	total := len(items)
	start := (page - 1) * perPage
	if start > total {
		start = total
	}
	end := min(start+perPage, total)

	data := struct {
		Items    []RelVM
		Page     int
		HasPrev  bool
		HasNext  bool
		PrevPage int
		NextPage int
	}{items[start:end], page, page > 1, end < total, page - 1, page + 1}

	renderFile(w, "web/templates/relations.html", data)
}
func main() {
	// API JSON
	http.HandleFunc("/api/artists", api.GetArtists)
	http.HandleFunc("/api/locations", api.GetLocations)
	http.HandleFunc("/api/dates", api.GetDates)
	http.HandleFunc("/api/relation", api.GetRelations)

	// Main page
	http.HandleFunc("/", artistsHandler) // home page: artists
	http.HandleFunc("/artists", artistsHandler)
	http.HandleFunc("/locations", locationsHandler)
	http.HandleFunc("/dates", datesHandler)
	http.HandleFunc("/relation", relationHandler)

	// static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	log.Println("🚀 Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
