package controllers

import (
	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"
	"github.com/tmohammad78/food-api/entity"
	"github.com/tmohammad78/food-api/services"
	"github.com/tmohammad78/food-api/validator"
	"strconv"
)

type FoodController interface {
	GetFood() []entity.Food
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	SaveFood(ctx *gin.Context) error
}

type controller struct {
	service services.FoodService
}

func (c *controller) Update(ctx *gin.Context) error {
	var food entity.Food
	err := ctx.ShouldBindJSON(&food)
	if err != nil {
		return err
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	food.ID = id
	err = validate.Struct(food)
	if err != nil {
		return err
	}
	c.service.Update(food)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var food entity.Food
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	food.ID = id
	c.service.Delete(food)
	return nil
}

var validate *validator2.Validate

func New(service services.FoodService) FoodController {
	validate = validator2.New()
	validate.RegisterValidation("is-cool", validator.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) GetFood() []entity.Food {
	return c.service.GetFood()
}

func (c *controller) SaveFood(ctx *gin.Context) error {
	var food entity.Food
	err := ctx.ShouldBindJSON(&food)
	if err != nil {
		return err
	}
	err = validate.Struct(food)
	if err != nil {
		return err
	}
	c.service.SaveFood(food)
	return nil
}
