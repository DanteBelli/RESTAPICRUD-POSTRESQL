package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Hola(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola , Prueba"))
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Hola)
	http.ListenAndServe(":3000", router)
}
