package app

import (
	"gotest/internal/resources/postgres"
	"gotest/internal/resources"
	"os"
)

func Run() {
	db := new(postgres.Database)
	defer db.Close()
	db.Init(os.Getenv("DATABASE_URL"))
	r := resources.Router{}
	r.Init()
	srv := resources.Server{}
	srv.Run("8080", r)
}