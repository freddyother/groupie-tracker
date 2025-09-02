package api

import (
	"encoding/json"
	"net/http"
)

type Date struct {
	ArtistID int      `json:"artistId"`
	Dates    []string `json:"dates"`
}

var dates = []Date{
	{ArtistID: 1, Dates: []string{"1963-03-22", "1965-08-15", "1969-01-30"}},
	{ArtistID: 2, Dates: []string{"1974-11-20", "1985-07-13", "1992-04-20"}},
	{ArtistID: 3, Dates: []string{"1967-05-12", "1973-03-01", "1994-10-29"}},
}

func GetDates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dates)
}
