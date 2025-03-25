package main

import (
	"net/http"

	"github.com/Dante-Belli/RESTAPICRUD-POSTRESQL/db"
	"github.com/Dante-Belli/RESTAPICRUD-POSTRESQL/models"
	"github.com/Dante-Belli/RESTAPICRUD-POSTRESQL/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})

	db.DB.AutoMigrate(models.User{})

	router := mux.NewRouter()

	router.HandleFunc("/", routes.Hola)

	http.ListenAndServe(":3000", router)
}
