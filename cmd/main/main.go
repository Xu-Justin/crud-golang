package main

import (
	"log"
	"fmt"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Xu-Justin/crud-golang/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	port := 8000
	fmt.Println("Serving application on port " + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), r))
}