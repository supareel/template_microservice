package main

import (
	"gomicro/database"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.ConnectToDB()
}