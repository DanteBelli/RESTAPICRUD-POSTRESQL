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

	router.HandleFunc("/users", routes.GetUsers).Methods("GET")
	router.HandleFunc("/tasks", routes.GetTasks).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUser).Methods("GET")
	router.HandleFunc("/tasks/{id}", routes.GetTask).Methods("GET")
	router.HandleFunc("/users", routes.PostUser).Methods("POST")
	router.HandleFunc("/tasks", routes.PostTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.DeleteTask).Methods("DELETE")
	router.HandleFunc("/users/{id}", routes.DeleteUser).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
