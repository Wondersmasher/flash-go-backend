package main

import (
	// "context"
	"fmt"
	"log"

	"github.com/flash-backend/config"
	"github.com/flash-backend/db"
	"github.com/flash-backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "github.com/redis/go-redis/v9"
	// "go.mongodb.org/mongo-driver/v2/mongo"
	// "go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	fmt.Println("Hello, Flash!")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	g.Use(gin.Logger())

	config.Env()
	db.InitRedis()
	db.InitMongoDB()

	routes.RegisterAllRoutes(g)

	g.Run(":8080")

}
