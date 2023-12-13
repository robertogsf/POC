package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robertogsf/POC/database"
	"github.com/robertogsf/POC/models"
	"github.com/robertogsf/POC/security"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	database.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	database.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	json.NewEncoder(w).Encode(&user)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var err error
	var result models.User
	json.NewDecoder(r.Body).Decode(&user)
	if err := database.DB.Where("email = ?", user.Email).First(&result).Error; err == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("El correo electrónico ya está registrado"))
		return
	}

	user = security.HashedPassword(user)

	createdUser := database.DB.Create(&user)
	err = createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var userbd models.User

	json.NewDecoder(r.Body).Decode(&user)

	database.DB.Where("email =?", user.Email).Find(&userbd)

	var is_correct bool = security.CheckPassword(user.Password, userbd.Password)

	if user.Email == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	} else if !is_correct {
		w.WriteHeader(401)
		w.Write([]byte("Password or Email is incorrect"))
		return
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	json.NewEncoder(w).Encode(security.GenerateAccessToken(userbd))
	log.Println(w.Header())
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	database.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	database.DB.Delete(&user)
	w.WriteHeader(http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Model(&user).Select("*").Updates(user)
	w.Write([]byte("Successful update"))
}
