package api

import "github.com/PCSB-Web-Team/FE-registration-portal-backend/config"

func NewServer() config.Initiators {
	app := config.NewApp()
	return app
}
