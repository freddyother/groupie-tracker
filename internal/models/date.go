package api

import (
	"encoding/json"
	"net/http"
)

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

var dates = []Date{
	{ID: 1, Dates: []string{"1963-03-22", "1965-07-15", "1969-01-30"}},
	{ID: 2, Dates: []string{"1974-04-12", "1977-10-07", "1986-07-12"}},
	{ID: 3, Dates: []string{"1967-08-05", "1973-03-01", "1994-10-20"}},
	{ID: 4, Dates: []string{"1969-01-12", "1971-11-08", "1979-07-24"}},
	{ID: 5, Dates: []string{"1964-04-16", "1972-06-09", "1981-12-18"}},
	{ID: 6, Dates: []string{"1989-06-15", "1991-09-24", "1993-12-13"}},
	{ID: 7, Dates: []string{"1993-02-22", "1997-05-21", "2007-10-10"}},
	{ID: 8, Dates: []string{"2000-07-10", "2002-08-26", "2016-12-02"}},
	{ID: 9, Dates: []string{"1980-03-10", "1987-09-21", "2017-07-01"}},
	{ID: 10, Dates: []string{"1983-07-25", "1991-08-12", "2016-11-18"}},
	{ID: 11, Dates: []string{"1975-04-17", "1980-07-25", "2008-11-09"}},
	{ID: 12, Dates: []string{"1967-01-04", "1970-06-30", "1971-12-12"}},
	{ID: 13, Dates: []string{"2006-01-23", "2009-03-24", "2018-05-11"}},
	{ID: 14, Dates: []string{"1995-07-04", "2005-06-14", "2021-02-05"}},
	{ID: 15, Dates: []string{"1987-07-21", "1992-05-12", "2019-06-21"}},
}

func GetDates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dates)
}
