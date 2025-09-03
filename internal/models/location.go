package api

import (
	"encoding/json"
	"net/http"
)

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

var locations = []Location{
	{ID: 1, Locations: []string{"Liverpool, UK", "London, UK", "Hamburg, Germany"}},
	{ID: 2, Locations: []string{"London, UK", "New York, USA", "Tokyo, Japan"}},
	{ID: 3, Locations: []string{"London, UK", "Los Angeles, USA", "Berlin, Germany"}},
	{ID: 4, Locations: []string{"London, UK", "New York, USA", "Paris, France"}},
	{ID: 5, Locations: []string{"London, UK", "Madrid, Spain", "Rome, Italy"}},
	{ID: 6, Locations: []string{"Seattle, USA", "New York, USA", "London, UK"}},
	{ID: 7, Locations: []string{"Oxford, UK", "Paris, France", "Berlin, Germany"}},
	{ID: 8, Locations: []string{"London, UK", "Paris, France", "Barcelona, Spain"}},
	{ID: 9, Locations: []string{"Dublin, Ireland", "Los Angeles, USA", "London, UK"}},
	{ID: 10, Locations: []string{"San Francisco, USA", "Berlin, Germany", "London, UK"}},
	{ID: 11, Locations: []string{"Sydney, Australia", "London, UK", "New York, USA"}},
	{ID: 12, Locations: []string{"Los Angeles, USA", "London, UK", "Amsterdam, Netherlands"}},
	{ID: 13, Locations: []string{"Sheffield, UK", "Paris, France", "Berlin, Germany"}},
	{ID: 14, Locations: []string{"Seattle, USA", "London, UK", "Madrid, Spain"}},
	{ID: 15, Locations: []string{"Los Angeles, USA", "London, UK", "Tokyo, Japan"}},
}

func GetLocations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locations)
}
