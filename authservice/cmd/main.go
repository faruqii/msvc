package main

import (
	"fmt"

	"github.com/faruqii/msvc/authservice/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	app.StartApplication()
}
