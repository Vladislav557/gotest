package app

import (
	"gotest/internal/resources"
	"gotest/internal/resources/postgres"
)

func Run() {
	postgres.Init("host=localhost port=6543 user=dev password=dev dbname=dev sslmode=disable")
	defer postgres.Close()
	r := resources.Router{}
	r.Init()
	srv := resources.Server{}
	srv.Run(":8000", r)
}
