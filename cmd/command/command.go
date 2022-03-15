package main

import (
	"log"

	"github.com/Renos-id/go-starter-template/cmd"
	"github.com/joho/godotenv"
)

func init() {
	// Init Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

}
func main() {
	cmd.Execute()
}
