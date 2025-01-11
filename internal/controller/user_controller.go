package controller

import (
	"encoding/json"
	"gotest/internal/model"
	"gotest/internal/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserController struct {
	us *service.UserService
}

func (c UserController) Update(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var u model.User
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}
	json.NewDecoder(req.Body).Decode(&u)
	c.us.Update(&u, id)
	json.NewEncoder(rw).Encode(u)
}

func (c UserController) Remove(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}
	c.us.Remove(id)
	json.NewEncoder(rw).Encode("success")
}

func (c UserController) Create(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var u model.User
	json.NewDecoder(req.Body).Decode(&u)
	c.us.Create(&u)
	json.NewEncoder(rw).Encode(u)
}

func (c UserController) GetUsers(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	users := c.us.GetUsers()
	json.NewEncoder(rw).Encode(users)
}

func (c UserController) GetUserByID(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}
	user := c.us.GetUserByID(id)
	if user.ID == 0 {
		rw.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(rw).Encode(user)
	}
}
