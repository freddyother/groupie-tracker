package api

import (
	"encoding/json"
	"net/http"
)

type Artist struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Image      string   `json:"image"`
	Year       int      `json:"year"`
	FirstAlbum string   `json:"firstAlbum"`
	Members    []string `json:"members"`
}

var artists = []Artist{
	{
		ID:         1,
		Name:       "The Beatles",
		Image:      "https://upload.wikimedia.org/wikipedia/en/2/2f/Beatles.png",
		Year:       1960,
		FirstAlbum: "Please Please Me (1963)",
		Members:    []string{"John Lennon", "Paul McCartney", "George Harrison", "Ringo Starr"},
	},
	{
		ID:         2,
		Name:       "Queen",
		Image:      "https://upload.wikimedia.org/wikipedia/en/0/05/Queen.png",
		Year:       1970,
		FirstAlbum: "Queen (1973)",
		Members:    []string{"Freddie Mercury", "Brian May", "Roger Taylor", "John Deacon"},
	},
	{
		ID:         3,
		Name:       "Pink Floyd",
		Image:      "https://upload.wikimedia.org/wikipedia/en/1/15/Pink_Floyd.png",
		Year:       1965,
		FirstAlbum: "The Piper at the Gates of Dawn (1967)",
		Members:    []string{"Syd Barrett", "Roger Waters", "Richard Wright", "Nick Mason", "David Gilmour"},
	},
}

func GetArtists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists)
}
