package routes

import (
	ctr "APIRest/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/", ctr.Home)
	r.HandleFunc("/api/models", ctr.AllPresidents).Methods("Get")
	r.HandleFunc("/api/models/{id}", ctr.ReturnID).Methods("Get")

	log.Fatal(http.ListenAndServe(":3333", r))
}
