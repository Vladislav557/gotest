package resources

import (
	"github.com/gorilla/mux"
	"gotest/internal/controller"
)

type Router struct {
	router *mux.Router
}

func (r *Router) Init() {
	r.router = mux.NewRouter()

	uc := controller.UserController{}

	r.router.HandleFunc("/users", uc.GetUsers).Methods("GET")
	r.router.HandleFunc("/users/{id}", uc.GetUserByID).Methods("GET")
	r.router.HandleFunc("/users", uc.Create).Methods("POST")
}