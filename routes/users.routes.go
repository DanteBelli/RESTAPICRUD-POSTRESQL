package routes

import (
	"encoding/json"
	"net/http"
	"github.com/Dante-Belli/RESTAPICRUD-POSTRESQL/db"
	"github.com/Dante-Belli/RESTAPICRUD-POSTRESQL/models"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params[]"id"])
	json.NewEncoder(w).Encode(&useuser)
	if  user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
}

func GetTask(w http.ResponseWriter, r *http.Request) {
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var user []models.User
	json.NewDecoder(r.Body).Decode(&user)
	createUser := db.DB.Create(&user)
	err := createUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
}

func PostTask(w http.ResponseWriter, r *http.Request) {
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var task []models.User
	params := mux.Vars()
	db.DB.First(&user,params["id"])
	if user.ID ==0{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	db.DB.Delete(&user)
}
