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
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}
	db.DB.Model(&user).Association("Task").Find(&user.Tasks)
	json.NewEncoder(w).Encode(&user)

}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	json.NewEncoder(w).Encode(&task)
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}
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
	var task []models.Task
	json.NewDecoder(r.Body).Decode(&task)
	createTask := db.DB.Create(&task)
	err := createTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var task []models.Task
	params := mux.Vars()
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not Found"))
		return
	}
	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user []models.User
	params := mux.Vars()
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
}
