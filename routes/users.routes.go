package routes

import (
	"net/http"

	"github.com/Dante-Belli/RESTAPICRUD-POSTRESQL/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

}
func GetUser(w http.ResponseWriter, r *http.Request) {
	var user []models.User

}
func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

}
func GetTask(w http.ResponseWriter, r *http.Request) {
	var task []models.User

}
func PostUser(w http.ResponseWriter, r *http.Request) {
	var user []models.User

}
func PostTask(w http.ResponseWriter, r *http.Request) {
	var task []models.Task

}
