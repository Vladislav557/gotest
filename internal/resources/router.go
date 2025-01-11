package resources

import (
	"encoding/json"
	"gotest/internal/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
}

func (r *Router) Init() {
	r.Router = mux.NewRouter()

	r.Router.HandleFunc("/health", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		res := map[string]string{"status": "ok"}
		jsonRes, err := json.Marshal(res)
		if err != nil {
			log.Fatal(err)
		}
		_, err = rw.Write(jsonRes)
		if err != nil {
			log.Fatal(err)
		}
	})

	uc := controller.UserController{}

	r.Router.HandleFunc("/users", uc.GetUsers).Methods("GET")
	r.Router.HandleFunc("/users", uc.Create).Methods("POST")
	r.Router.HandleFunc("/users/{id}", uc.GetUserByID).Methods("GET")
	r.Router.HandleFunc("/users/{id}", uc.Remove).Methods("DELETE")
	r.Router.HandleFunc("/users/{id}", uc.Update).Methods("PUT")
}
