package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tmohammad78/food-api/entity"
	"github.com/tmohammad78/food-api/services"
)

type FoodController interface {
	GetFood() []entity.Food
	SaveFood(ctx *gin.Context) entity.Food
}

type controller struct {
	service services.FoodService
}

func New(service services.FoodService) FoodController {
	return &controller{
		service: service,
	}
}

func (c *controller) GetFood() []entity.Food {
	return c.service.GetFood()
}

func (c *controller) SaveFood(ctx *gin.Context) entity.Food {
	var food entity.Food
	ctx.BindJSON(&food)
	c.service.SaveFood(food)
	return food
}
