package routes

import "net/http"

func Hola(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola , Prueba"))
}
