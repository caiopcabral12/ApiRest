package controller

import (
	"APIRest/database"
	md "APIRest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")

}

func AllPresidents(w http.ResponseWriter, r *http.Request) {
	var p []md.Presidents
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)

}

func ReturnID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p md.Presidents
	database.DB.First(&p, id)

	json.NewEncoder(w).Encode(p)

}

func NewPresidents(w http.ResponseWriter, r *http.Request) {
	var newPresident md.Presidents
	json.NewDecoder(r.Body).Decode(&newPresident)
	database.DB.Create(&newPresident)
	json.NewEncoder(w).Encode(newPresident)

}

func DeletePresidents(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var deletePresident md.Presidents
	database.DB.Delete(&deletePresident, id)
	json.NewEncoder(w).Encode(deletePresident)
}

func UpdatePresidents(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updatePresident md.Presidents
	database.DB.First(&updatePresident, id)
	json.NewDecoder(r.Body).Decode(&updatePresident)
	database.DB.Save(&updatePresident)
	json.NewEncoder(w).Encode(updatePresident)
}
