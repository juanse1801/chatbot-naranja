package main

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/juanse1801/chatbot-naranja/cmd/server/routes"
	"github.com/juanse1801/chatbot-naranja/pkg/configs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Set and start Cron Scheduler
	cron := gocron.NewScheduler(time.Local)
	cron.StartAsync()

	// Load .env file
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	db, errDB := configs.ConnectDB()
	if errDB != nil {
		panic(errDB)
	}

	r := gin.Default()

	router := routes.NewRouter(r, cron, db)
	router.MapRoutes()

	// Run application
	if err := r.Run(); err != nil {
		panic(err)
	}

}
