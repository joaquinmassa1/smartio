package routes

import "net/http"

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
