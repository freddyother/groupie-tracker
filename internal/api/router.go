package api

import "net/http"

// Aquí se definen las rutas
func RegisterRoutes() {
	http.HandleFunc("/hello", HelloHandler)
}
