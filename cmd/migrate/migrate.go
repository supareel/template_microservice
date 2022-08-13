package main

import (
	"gomicro/config"
	"gomicro/database"
	"gomicro/pkg"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.development")
	config.LoadEnvConfig()
	// connect to DB
	db, err := database.NewPostgresClient(config.EnvConfig.DB_USERNAME, config.EnvConfig.DB_PASS, config.EnvConfig.DB_SERVER, config.EnvConfig.DB_NAME, config.EnvConfig.DB_SSL, config.EnvConfig.DB_PORT, 30)
	if err != nil {
		pkg.FancyHandleError(err)
	}
	db.MigrateDB(40)
	defer db.CloseClient()
}
