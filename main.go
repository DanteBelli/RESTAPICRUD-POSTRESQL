package main

import (
	"net/http"

	"github.com/Dante-Belli/RESTAPICRUD-POSTRESQL/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", routes.Hola)
	http.ListenAndServe(":3000", router)
}
