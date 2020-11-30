package services

import (
	"github.com/tmohammad78/food-api/entity"
	"github.com/tmohammad78/food-api/repository"
)

type FoodService interface {
	GetFood() []entity.Food
	Update(food entity.Food)
	Delete(food entity.Food)
	SaveFood(food entity.Food) entity.Food
}

type foodService struct {
	foodRepository repository.FoodRepository
}

func (service *foodService) Update(food entity.Food) {
	service.foodRepository.Update(food)
}

func (service *foodService) Delete(food entity.Food) {
	service.foodRepository.Delete(food)
}

func New(repo repository.FoodRepository) FoodService {
	return &foodService{
		foodRepository: repo,
	}
}

func (service *foodService) GetFood() []entity.Food {
	return service.foodRepository.FindAll()
}

func (service *foodService) SaveFood(food entity.Food) entity.Food {
	service.foodRepository.Save(food)
	return food
}
