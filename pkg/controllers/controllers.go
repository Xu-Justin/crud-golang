package controllers

import (
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/Xu-Justin/crud-golang/pkg/utils"
	"github.com/Xu-Justin/crud-golang/pkg/models"
)

func Get(w http.ResponseWriter, r *http.Request) {
	users := models.Get()
	
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 0, 0)
	user := models.GetById(id)

	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Create(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	utils.ParseBody(r, user)
	user = models.Create(user)

	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 0, 0)
	user := &models.User{}
	utils.ParseBody(r, user)
	
	var newUser *models.User
	if user.Name != "" { newUser = models.UpdateName(id, user.Name) }
	if user.Age != 0 { newUser = models.UpdateAge(id, user.Age) }

	res, _ := json.Marshal(newUser)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 0, 0)
	user := models.Delete(id)

	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}