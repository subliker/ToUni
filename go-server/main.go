package main

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/subliker/ToUni/go-server/db"
	"github.com/subliker/ToUni/go-server/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dataBase := db.DataBase{}
	dataBase.Init()

	var router router.Router
	router.SetDataBase(&dataBase)
	router.SetupRouter()
	router.Run(os.Getenv("SERVER_PORT"))
}
