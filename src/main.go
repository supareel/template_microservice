package main

import (
	"flag"
	"fmt"
	"gomicro/docs"
	"gomicro/src/config"
	"gomicro/src/database"
	"gomicro/src/internal/taskmanager"
	"gomicro/src/pkg"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title									Swagger Example API
// @version								1.0
// @description							This is a sample server celler server.
// @termsOfService							http://swagger.io/terms/
//
// @contact.name							API Support
// @contact.url							http://www.swagger.io/support
// @contact.email							support@swagger.io
//
// @license.name							Apache 2.0
// @license.url							http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host									localhost:8080
// @BasePath								/api/v1
//
// @securitydefinitions.oauth2.accessCode	OAuth2AccessCode
// @tokenUrl								https://example.com/oauth/token
// @authorizationUrl						https://example.com/oauth/authorize
// @scope.admin							Grants read and write access to administrative information
//
// @externalDocs.description				OpenAPI
// @externalDocs.url						https://swagger.io/resources/open-api/
func main() {

	var (
		migrateFlag bool
	)

	flag.BoolVar(&migrateFlag, "migrate", false, "used to run migration command")

	flag.Parse()

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Task manager APIs"
	docs.SwaggerInfo.Description = "A sample OpenApi docs generator using swag"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	godotenv.Load(".env.development")
	config.LoadEnvConfig()

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

	if migrateFlag {
		db.MigrateDB(40)
	}

	router := gin.Default()

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// version router
	// v1 := router.Group("/api/v1"){}

	// repository
	repo := taskmanager.NewPostgresRepository(db)
	taskApi := taskmanager.NewTaskApi(repo)
	taskmanager.RegisterHandlers(router, taskApi)

	// start server
	httpEndpoint := fmt.Sprintf(":%d", config.EnvConfig.HTTP_PORT)
	if err = router.Run(httpEndpoint); err != nil {
		log.Fatal("Cannot Start HTTP Server : ", err)
	}
}
