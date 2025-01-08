package resources

import (
	"log"
	"net/http"
)

type Server struct{}

func (s *Server) Run(p string, r Router) {
	log.Fatal(http.ListenAndServe(p, r.router))
}