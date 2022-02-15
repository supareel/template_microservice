package main

import (
	"gomicro/database"
	"gomicro/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.development")
	database.ConnectToDB()

	//tokenMaker, err := token.NewPasetoMaker(os.Getenv("TOKEN_SYMMETRIC_KEY"))
	// if err != nil {
	//   log.Fatal(err)
	// }

	mux := gin.Default()

	// define routes
	routes.LoginRoutes(mux)

	// start server
	mux.Run()
}
