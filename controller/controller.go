package controller

import (
	md "APIRest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")

}

func AllPresidents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(md.President)

}

func ReturnID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, Presidents := range md.President {
		if strconv.Itoa(Presidents.Id) == id {
			json.NewEncoder(w).Encode(Presidents)

		}
	}
}
