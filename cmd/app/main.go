package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"restApi/internal/app/config"
)

func init() {
	godotenv.Load()
	config.InitLog()
}

func main() {
	port := os.Getenv("PORT")
	app := config.SetupRouter()

	err := app.Run(":" + port)
	if err != nil {
		log.Error("Error running server: ", err)
	}
}
