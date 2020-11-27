package services

import "github.com/tmohammad78/food-api/entity"

type FoodService interface {
	GetFood() []entity.Food
	SaveFood(food entity.Food) entity.Food
}

type foodService struct {
	foods []entity.Food
}

func New() FoodService {
	return &foodService{}
}

func (service *foodService) GetFood() []entity.Food {
	return service.foods
}

func (service *foodService) SaveFood(food entity.Food) entity.Food {
	service.foods = append(service.foods, food)
	return food
}
