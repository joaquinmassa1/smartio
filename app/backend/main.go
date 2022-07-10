package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joaquinmassa1/smartio/routes"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", routes.Home)

	http.ListenAndServe(":3000", r)
}
