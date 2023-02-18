package main

import (
	"log"

	"github.com/juanse1801/chatbot-naranja/cmd/server/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	router := routes.NewRouter(r)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}

}
