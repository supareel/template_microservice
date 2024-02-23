package main

import (
	"fmt"
	"gomicro/src/config"
	"gomicro/src/database"
	"gomicro/src/internal/taskmanager"
	taskmanager_gen "gomicro/src/internal/taskmanager/_generated_"
	"gomicro/src/pkg"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.development")
	config.LoadEnvConfig()

	router := gin.Default()

	// connect to DB
	db, err := database.NewPostgresClient(
		config.EnvConfig.DB_USERNAME,
		config.EnvConfig.DB_PASS,
		config.EnvConfig.DB_SERVER,
		config.EnvConfig.DB_NAME,
		config.EnvConfig.DB_SSL,
		config.EnvConfig.DB_PORT, 30)
	if err != nil {
		pkg.FancyHandleError(err)
	}
	defer db.CloseClient()

	// repository
	repo := taskmanager.NewPostgresRepository(db)
	taskApi := taskmanager.NewTaskApi(repo)
	taskmanager_gen.RegisterHandlers(router, taskApi)

	// start server
	httpEndpoint := fmt.Sprintf(":%d", config.EnvConfig.HTTP_PORT)
	if err = router.Run(httpEndpoint); err != nil {
		log.Fatal("Cannot Start HTTP Server : ", err)
	}
}
