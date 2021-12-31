package main

import (
	"fmt"
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

	allAccounts := database.DeleteAccountBalance(4)
	fmt.Println(allAccounts)
}
