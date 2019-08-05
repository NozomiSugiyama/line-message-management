package main

import (
	"log"
	"os"

	"time"

	"github.com/jinzhu/gorm"

	"api/database"
	"api/database/model"
	"api/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	var mode string
	if len(os.Args) >= 2 {
		mode = os.Args[1]
	}
	var envFile string
	if mode == "dev" {
		envFile = ".env.development"
	} else {
		envFile = ".env.production"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Println("Failed load " + envFile + " file")
	}

	// Get Environments
	port := os.Getenv("PORT")
	lineChannelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	lineChannelAccessToken := os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")
	providerWebOrigin := os.Getenv("PROVIDER_WEB_ORIGIN")
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresDatabase := os.Getenv("POSTGRES_DATABASE")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresSSLMode := os.Getenv("POSTGRES_SSL_MODE")

	// default port number
	if port == "" {
		port = "8080"
	}

	var db *gorm.DB

	// Connection attempt when connecting to DB with docker-compose
	for i := 0; i < 10; i++ {
		db, err = database.Initialize(postgresHost, postgresPort, postgresUser, postgresDatabase, postgresPassword, postgresSSLMode)
		if err != nil {
			log.Println("Failed database initialize")
			log.Println(err)
			if i >= 9 {
				panic(err)
			}
			time.Sleep(10 * time.Second)
		} else {
			log.Println("Connected database")
			break
		}
	}

	defer db.Close()

	userRepo := model.NewUserRepository(db)
	nonceRepo := model.NewNonceRepository(db)
	lineUserRepo := model.NewLineUserRepository(db)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())
	router.Use(database.Inject(db))

	userHandler := handler.NewUserHandler(userRepo)
	router.GET("/users", userHandler.GetUsers)
	router.GET("/users/:userid", userHandler.GetUserByID)

	authHandler := handler.NewAuthHandler(userRepo, nonceRepo)
	router.POST("/auth/client-sign-in", authHandler.ClientSignIn)

	hookHandler := handler.NewHookHandler(userRepo, nonceRepo, lineUserRepo, lineChannelSecret, lineChannelAccessToken, providerWebOrigin)
	router.POST("/hook", hookHandler.PostHook)

	router.Run(":" + port)
	log.Println("Service starting on port " + port)
}
