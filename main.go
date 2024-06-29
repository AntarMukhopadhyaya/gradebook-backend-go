package main

import (
	"github.com/gin-gonic/gin"
	"gradebook-api/controllers"
	"gradebook-api/helpers"
	"gradebook-api/middleware"
)

func init() {
	helpers.LoadEnvVariables()
	helpers.ConnectToDb()
	helpers.MigrateDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/welcome", middleware.RequireAuth, controllers.Validate)
	r.POST("/u/register", controllers.SignUp)
	r.POST("/u/login", controllers.Login)
	r.Run()
}
