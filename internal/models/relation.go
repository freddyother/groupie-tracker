package api

import (
	"encoding/json"
	"net/http"
)

type Relation struct {
	ArtistID int      `json:"artistId"`
	Cities   []string `json:"cities"`
	Dates    []string `json:"dates"`
}

var relations = []Relation{
	{
		ArtistID: 1,
		Cities:   []string{"Liverpool", "New York"},
		Dates:    []string{"1963-03-22", "1965-08-15"},
	},
	{
		ArtistID: 2,
		Cities:   []string{"London", "Paris"},
		Dates:    []string{"1974-11-20", "1985-07-13"},
	},
	{
		ArtistID: 3,
		Cities:   []string{"Cambridge", "Los Angeles"},
		Dates:    []string{"1967-05-12", "1973-03-01"},
	},
}

func GetRelations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(relations)
}
