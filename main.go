package main

import (
	"os"

	"github.com/PCSB-Web-Team/FE-registration-portal-backend/api"
)

func init() {
	server := api.NewServer()
	server.Start(os.Getenv("PORT"))
}
