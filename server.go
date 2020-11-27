package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tmohammad78/food-api/controllers"
	"github.com/tmohammad78/food-api/services"
)

var (
	foodService    services.FoodService       = services.New()
	foodController controllers.FoodController = controllers.New(foodService)
)

func main() {
	server := gin.Default()
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
