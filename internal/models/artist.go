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
		Image:      "static/img/the_beatles.jpg",
		Year:       1960,
		FirstAlbum: "Please Please Me (1963)",
		Members:    []string{"John Lennon", "Paul McCartney", "George Harrison", "Ringo Starr"},
	},
	{
		ID:         2,
		Name:       "Queen",
		Image:      "static/img/queen.jpg",
		Year:       1970,
		FirstAlbum: "Queen (1973)",
		Members:    []string{"Freddie Mercury", "Brian May", "Roger Taylor", "John Deacon"},
	},
	{
		ID:         3,
		Name:       "Pink Floyd",
		Image:      "static/img/pink_floyd.jpg",
		Year:       1965,
		FirstAlbum: "The Piper at the Gates of Dawn (1967)",
		Members:    []string{"Syd Barrett", "Roger Waters", "Richard Wright", "Nick Mason", "David Gilmour"},
	},
	{
		ID:         4,
		Name:       "Led Zeppelin",
		Image:      "static/img/led_zeppelin.jpg",
		Year:       1968,
		FirstAlbum: "Led Zeppelin (1969)",
		Members:    []string{"Robert Plant", "Jimmy Page", "John Paul Jones", "John Bonham"},
	},
	{
		ID:         5,
		Name:       "The Rolling Stones",
		Image:      "static/img/the_rolling_stones.jpg",
		Year:       1962,
		FirstAlbum: "The Rolling Stones (1964)",
		Members:    []string{"Mick Jagger", "Keith Richards", "Charlie Watts", "Ronnie Wood"},
	},
	{
		ID:         6,
		Name:       "Nirvana",
		Image:      "static/img/nirvana.jpg",
		Year:       1987,
		FirstAlbum: "Bleach (1989)",
		Members:    []string{"Kurt Cobain", "Krist Novoselic", "Dave Grohl"},
	},
	{
		ID:         7,
		Name:       "Radiohead",
		Image:      "static/img/radiohead.jpg",
		Year:       1985,
		FirstAlbum: "Pablo Honey (1993)",
		Members:    []string{"Thom Yorke", "Jonny Greenwood", "Ed O'Brien", "Colin Greenwood", "Philip Selway"},
	},
	{
		ID:         8,
		Name:       "Coldplay",
		Image:      "static/img/coldplay.jpg",
		Year:       1996,
		FirstAlbum: "Parachutes (2000)",
		Members:    []string{"Chris Martin", "Jonny Buckland", "Guy Berryman", "Will Champion"},
	},
	{
		ID:         9,
		Name:       "U2",
		Image:      "static/img/u2.jpg",
		Year:       1976,
		FirstAlbum: "Boy (1980)",
		Members:    []string{"Bono", "The Edge", "Adam Clayton", "Larry Mullen Jr."},
	},
	{
		ID:         10,
		Name:       "Metallica",
		Image:      "static/img/metallica.jpg",
		Year:       1981,
		FirstAlbum: "Kill 'Em All (1983)",
		Members:    []string{"James Hetfield", "Lars Ulrich", "Kirk Hammett", "Robert Trujillo"},
	},
	{
		ID:         11,
		Name:       "AC/DC",
		Image:      "static/img/ac_dc.jpg",
		Year:       1973,
		FirstAlbum: "High Voltage (1975)",
		Members:    []string{"Angus Young", "Malcolm Young", "Bon Scott", "Brian Johnson", "Phil Rudd"},
	},
	{
		ID:         12,
		Name:       "The Doors",
		Image:      "static/img/the_doors.jpg",
		Year:       1965,
		FirstAlbum: "The Doors (1967)",
		Members:    []string{"Jim Morrison", "Robby Krieger", "Ray Manzarek", "John Densmore"},
	},
	{
		ID:         13,
		Name:       "Arctic Monkeys",
		Image:      "static/img/arctic_monkeys.jpg",
		Year:       2002,
		FirstAlbum: "Whatever People Say I Am, That's What I'm Not (2006)",
		Members:    []string{"Alex Turner", "Jamie Cook", "Nick O'Malley", "Matt Helders"},
	},
	{
		ID:         14,
		Name:       "Foo Fighters",
		Image:      "static/img/foo_fighters.jpg",
		Year:       1994,
		FirstAlbum: "Foo Fighters (1995)",
		Members:    []string{"Dave Grohl", "Nate Mendel", "Pat Smear", "Taylor Hawkins", "Chris Shiflett"},
	},
	{
		ID:         15,
		Name:       "Guns N' Roses",
		Image:      "static/img/guns_n_roses.jpg",
		Year:       1985,
		FirstAlbum: "Appetite for Destruction (1987)",
		Members:    []string{"Axl Rose", "Slash", "Duff McKagan", "Izzy Stradlin", "Steven Adler"},
	},
}

func GetArtists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists)
}

func GetArtistsData() []Artist {
	return artists
}
