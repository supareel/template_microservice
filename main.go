package main

import (
	"gomicro/database"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.ConnectToDB()
	// database.CreateAccount(ent.Accounts{
	// 	Owner: "sourabh",
	// 	Balance: 100,
	// 	Currency: "INR",
	// })
}
