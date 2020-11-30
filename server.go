package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tmohammad78/food-api/controllers"
	"github.com/tmohammad78/food-api/middlewares"
	"github.com/tmohammad78/food-api/repository"
	"github.com/tmohammad78/food-api/services"
	"io"
	"net/http"
	"os"
)

var (
	foodRepository  repository.FoodRepository   = repository.NewFoodRepository()
	foodService     services.FoodService        = services.New(foodRepository)
	loginService    services.LoginService       = services.NewLoginService()
	jwtService      services.JwtService         = services.NewJWTService()
	foodController  controllers.FoodController  = controllers.New(foodService)
	loginController controllers.LoginController = controllers.NewLoginController(loginService, jwtService)
)

/** method for save logger files */
func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	defer foodRepository.CloseDB()
	setupLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())
	server.POST("/login", func(context *gin.Context) {
		token := loginController.Login(context)
		if token != "" {
			context.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			context.JSON(http.StatusUnauthorized, nil)
		}
	})
	apiGroup := server.Group("/api", middlewares.AuthorizedJWT())
	{
		apiGroup.GET("/food", func(context *gin.Context) {
			context.JSON(200, foodController.GetFood())
		})
		apiGroup.POST("/food", func(context *gin.Context) {
			err := foodController.SaveFood(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "input is valid"})
			}

		})

		apiGroup.PUT("/food/:id", func(context *gin.Context) {
			err := foodController.Update(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "input is valid"})
			}
		})
		apiGroup.DELETE("/food/:id", func(context *gin.Context) {
			err := foodController.Delete(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "input is valid"})
			}
		})
	}

	server.Run()
}
