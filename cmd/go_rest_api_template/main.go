package main

import (
	"log"
	"rest_api_template/internal/app/api"
	"rest_api_template/internal/app/models"

	"gopkg.in/ini.v1"
)

func main() {

	cfg, err := ini.Load("configs/config.ini")

	if err != nil {
		log.Fatal(err)
	}

	db := models.InitDatabase(cfg)
	defer db.Close()

	server := api.NewServer(cfg)

	server.Logger().Infof("Start api server on %s", cfg.Section("server").Key("listen_addr").String())

	if err := server.Start(); err != nil {
		server.Logger().Fatal(err)
	}

}
