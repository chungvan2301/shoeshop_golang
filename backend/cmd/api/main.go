package main

import (
	"log"

	"github.com/chungvan2301/shoeshop/backend/pkg/di"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the env file")
	}

	di.InitializeApp()
}
