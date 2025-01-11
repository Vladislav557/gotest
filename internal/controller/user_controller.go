package controller

import (
	"encoding/json"
	"go.uber.org/zap"
	"gotest/internal/model"
	"gotest/internal/service"
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
		zap.S().Error(err)
		return
	}
	err = json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		zap.S().Error(err)
		return
	}
	c.us.Update(&u, id)
	err = json.NewEncoder(rw).Encode(u)
	if err != nil {
		zap.S().Error(err)
		return
	}
}

func (c UserController) Remove(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		zap.S().Error(err)
		return
	}
	c.us.Remove(id)
	err = json.NewEncoder(rw).Encode("success")
	if err != nil {
		zap.S().Error(err)
		return
	}
}

func (c UserController) Create(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var u model.User
	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		zap.S().Error(err)
		return
	}
	c.us.Create(&u)
	err = json.NewEncoder(rw).Encode(u)
	if err != nil {
		zap.S().Error(err)
		return
	}
}

func (c UserController) GetUsers(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	users := c.us.GetUsers()
	err := json.NewEncoder(rw).Encode(users)
	if err != nil {
		zap.S().Error(err)
		return
	}
}

func (c UserController) GetUserByID(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		zap.S().Error(err)
		return
	}
	user := c.us.GetUserByID(id)
	if user.ID == 0 {
		rw.WriteHeader(http.StatusNotFound)
	} else {
		err = json.NewEncoder(rw).Encode(user)
		if err != nil {
			zap.S().Error(err)
			return
		}
	}
}
