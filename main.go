package main

import (
	"ExpenseAPI/controllers"
	"ExpenseAPI/initializers"
	"ExpenseAPI/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {

	router := gin.Default()

	router.POST("/register", controllers.Signup)

	router.POST("/login", controllers.Login)

	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	router.POST("/expenses/add", middleware.RequireAuth, controllers.AddExpense)

	err := router.Run("localhost:8080")
	if err != nil {
		panic("Failed to start server")
	}
}
