package routes

import (
	ctr "APIRest/controller"
	mw "APIRest/middleaware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(mw.Middleaware)

	r.HandleFunc("/", ctr.Home)
	r.HandleFunc("/api/models", ctr.AllPresidents).Methods("Get")
	r.HandleFunc("/api/models/{id}", ctr.ReturnID).Methods("Get")
	r.HandleFunc("/api/models", ctr.NewPresidents).Methods("Post")
	r.HandleFunc("/api/models/{id}", ctr.DeletePresidents).Methods("Delete")
	r.HandleFunc("/api/models/{id}", ctr.UpdatePresidents).Methods("Put")

	log.Fatal(http.ListenAndServe(":3333", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
