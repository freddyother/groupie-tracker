package api

import "net/http"

// Ejemplo de handler
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola desde Groupie Tracker"))
}
