package main

import (
	"fmt"
	"github.com/anthcor/gin-sampler/cmd/gin-sampler/apis"
	"github.com/anthcor/gin-sampler/cmd/gin-sampler/config"
	"github.com/anthcor/gin-sampler/cmd/gin-sampler/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // the postgres driver
	"log"
	"os"
)

func main() {
	os.Setenv("TZ", "UTC")

	// load application configurations
	if err := config.LoadConfig(); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	byUserID := "/users/:id"
	v1 := r.Group("/api/v1")
	{
		v1.GET(byUserID, apis.GetUser)
		v1.GET("/users", apis.GetUsers)
		v1.POST("/users", apis.CreateUser)
		v1.PATCH(byUserID, apis.UpdateUser)
		v1.DELETE(byUserID, apis.DeleteUser)
	}

	r.GET("/health", apis.Healthcheck)

	config.Config.DB, config.Config.DBErr = gorm.Open("postgres", config.Config.DSN)
	if config.Config.DBErr != nil {
		panic(config.Config.DBErr)
	}

	// This is needed for generation of schema for postgres image
	config.Config.DB.AutoMigrate(&models.User{})

	defer config.Config.DB.Close()

	log.Println("Successfully connected to database")

	r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))
}
