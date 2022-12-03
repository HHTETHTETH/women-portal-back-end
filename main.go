package main

import (
	"log"

	router "back/routers"
	driver "back/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"os"
)

const (
	envAPIMongoURL = "API_MONGO_URL"
	envAPIPort     = "API_PORT"
	envAPIMode     = "API_MODE"
)

func main() {
	goPath := os.Getenv("GOPATH")
	filePath := goPath + "/src/back/.env-example"
	err := godotenv.Load(filePath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gin.SetMode(os.Getenv(envAPIMode))

	driver.InitDriver(os.Getenv(envAPIMongoURL))
	router.InitConnection(os.Getenv(envAPIPort))
}
