package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tmohammad78/food-api/controllers"
	"github.com/tmohammad78/food-api/middlewares"
	"github.com/tmohammad78/food-api/services"
	"io"
	"os"
)

var (
	foodService    services.FoodService       = services.New()
	foodController controllers.FoodController = controllers.New(foodService)
)

/** method for save logger files */
func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	setupLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())
	server.GET("/food", func(context *gin.Context) {
		context.JSON(200, foodController.GetFood())
	})
	server.POST("/food", func(context *gin.Context) {
		context.JSON(200, foodController.SaveFood(context))
	})
	//server.GET("/food", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"message": "ok",
	//	})
	//})
	server.Run()
}
