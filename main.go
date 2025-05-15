package main

import (
	"log"
	"url-shortener/routers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := routers.SetupRouter()
	r.Run(":9000")
}
