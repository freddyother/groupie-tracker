package api

import (
	"encoding/json"
	"net/http"
)

type Relation struct {
	ID        int                 `json:"id"`
	Relations map[string][]string `json:"relations"`
}

var relations = []Relation{
	{ID: 1, Relations: map[string][]string{"Liverpool, UK": {"1963-03-22"}, "London, UK": {"1965-07-15"}, "Hamburg, Germany": {"1969-01-30"}}},
	{ID: 2, Relations: map[string][]string{"London, UK": {"1974-04-12"}, "New York, USA": {"1977-10-07"}, "Tokyo, Japan": {"1986-07-12"}}},
	{ID: 3, Relations: map[string][]string{"London, UK": {"1967-08-05"}, "Los Angeles, USA": {"1973-03-01"}, "Berlin, Germany": {"1994-10-20"}}},
	{ID: 4, Relations: map[string][]string{"London, UK": {"1969-01-12"}, "New York, USA": {"1971-11-08"}, "Paris, France": {"1979-07-24"}}},
	{ID: 5, Relations: map[string][]string{"London, UK": {"1964-04-16"}, "Madrid, Spain": {"1972-06-09"}, "Rome, Italy": {"1981-12-18"}}},
	{ID: 6, Relations: map[string][]string{"Seattle, USA": {"1989-06-15"}, "New York, USA": {"1991-09-24"}, "London, UK": {"1993-12-13"}}},
	{ID: 7, Relations: map[string][]string{"Oxford, UK": {"1993-02-22"}, "Paris, France": {"1997-05-21"}, "Berlin, Germany": {"2007-10-10"}}},
	{ID: 8, Relations: map[string][]string{"London, UK": {"2000-07-10"}, "Paris, France": {"2002-08-26"}, "Barcelona, Spain": {"2016-12-02"}}},
	{ID: 9, Relations: map[string][]string{"Dublin, Ireland": {"1980-03-10"}, "Los Angeles, USA": {"1987-09-21"}, "London, UK": {"2017-07-01"}}},
	{ID: 10, Relations: map[string][]string{"San Francisco, USA": {"1983-07-25"}, "Berlin, Germany": {"1991-08-12"}, "London, UK": {"2016-11-18"}}},
	{ID: 11, Relations: map[string][]string{"Sydney, Australia": {"1975-04-17"}, "London, UK": {"1980-07-25"}, "New York, USA": {"2008-11-09"}}},
	{ID: 12, Relations: map[string][]string{"Los Angeles, USA": {"1967-01-04"}, "London, UK": {"1970-06-30"}, "Amsterdam, Netherlands": {"1971-12-12"}}},
	{ID: 13, Relations: map[string][]string{"Sheffield, UK": {"2006-01-23"}, "Paris, France": {"2009-03-24"}, "Berlin, Germany": {"2018-05-11"}}},
	{ID: 14, Relations: map[string][]string{"Seattle, USA": {"1995-07-04"}, "London, UK": {"2005-06-14"}, "Madrid, Spain": {"2021-02-05"}}},
	{ID: 15, Relations: map[string][]string{"Los Angeles, USA": {"1987-07-21"}, "London, UK": {"1992-05-12"}, "Tokyo, Japan": {"2019-06-21"}}},
}

func GetRelations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(relations)
}
func GetRelationsData() []Relation { return relations }
