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
	r.POST("/a/create", middleware.RequireAuth, controllers.CreateAssignment)
	r.GET("/show-assignment", middleware.RequireAuth, controllers.ShowAssignment)
	r.GET("/a/all", middleware.RequireAuth, controllers.IndexAssignment)
	r.Run()
}
