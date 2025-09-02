package api

import (
	"encoding/json"
	"net/http"
)

type Location struct {
	ArtistID int      `json:"artistId"`
	Cities   []string `json:"cities"`
}

var locations = []Location{
	{ArtistID: 1, Cities: []string{"Liverpool", "Hamburg", "New York", "Tokyo"}},
	{ArtistID: 2, Cities: []string{"London", "Paris", "Rio de Janeiro", "Berlin"}},
	{ArtistID: 3, Cities: []string{"Cambridge", "Los Angeles", "Moscow", "Sydney"}},
}

func GetLocations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locations)
}
