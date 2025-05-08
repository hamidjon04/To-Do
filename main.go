package main

import (
	"fmt"
	"todo/api"
	"todo/config"
	postgres "todo/storage"
)

func main() {
	cfg := config.LoadConfig()

	db, err := postgres.Connect(cfg)
	if err != nil {
		fmt.Println("Error connect to postgres: ", err)
	}
	defer db.Close()

	storage := postgres.Storage(db)
	router := api.Router(storage)
	router.Run(cfg.TO_DO)
}
