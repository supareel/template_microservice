package main

import (
	"gomicro/config"
	"gomicro/database"
	"gomicro/domain/taskmanager"
	"gomicro/pkg"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func handleArgs(args []string, db database.Database) {
	if args[1] == "-migrate" {
		db.MigrateDB(40)
	}
}

func main() {
	godotenv.Load(".env.development")
	config.LoadEnvConfig()
	
	if config.EnvConfig.GIN_MODE == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	// router
	router := gin.Default()

	// connect to DB
	db, err := database.NewPostgresClient(config.EnvConfig.DB_USERNAME, config.EnvConfig.DB_PASS, config.EnvConfig.DB_SERVER, config.EnvConfig.DB_NAME, config.EnvConfig.DB_SSL, config.EnvConfig.DB_PORT, 30)
	if err != nil {
		pkg.FancyHandleError(err)
	}

	if len(os.Args) > 1 {
		handleArgs(os.Args, db)
		return
	}
	// repository
	repo := taskmanager.NewPostgresRepository(db)
	svc := taskmanager.NewTaskService(repo)
	taskRouter := taskmanager.NewTaskHandler(svc)

	// define routes
	taskRouter.Register(router)
	// rest.NewTaskHandler.Register(router)

	// start server
	router.Run(":" + config.EnvConfig.APP_PORT)
}
