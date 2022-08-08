package routes

import (
	"github.com/gorilla/mux"
	"github.com/Xu-Justin/crud-golang/pkg/controllers"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.Get).Methods("GET")
	router.HandleFunc("/{id}", controllers.GetById).Methods("GET")
	router.HandleFunc("/", controllers.Create).Methods("POST")
	router.HandleFunc("/{id}", controllers.Update).Methods("PUT")
	router.HandleFunc("/{id}", controllers.Delete).Methods("DELETE")
}