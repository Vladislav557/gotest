package resources

import (
	"encoding/json"
	"gotest/internal/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router
}

func (r *Router) Init() {
	r.router = mux.NewRouter()

	r.router.HandleFunc("/health", func (rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		res := map[string]string{"status": "ok"}
		json, err := json.Marshal(res)
		if err != nil {
			log.Fatal(err)
		}
		rw.Write(json)
	})

	uc := controller.UserController{}

	r.router.HandleFunc("/users", uc.GetUsers).Methods("GET")
	r.router.HandleFunc("/users", uc.Create).Methods("POST")
	r.router.HandleFunc("/users/{id}", uc.GetUserByID).Methods("GET")
	r.router.HandleFunc("/users/{id}", uc.Remove).Methods("DELETE")
	r.router.HandleFunc("/users/{id}", uc.Update).Methods("PUT")
}