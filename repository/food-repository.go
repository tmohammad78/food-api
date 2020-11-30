package repository

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tmohammad78/food-api/entity"
	"gorm.io/gorm"
)

type FoodRepository interface {
	Save(food entity.Food)
	Update(food entity.Food)
	Delete(food entity.Food)
	FindAll() []entity.Food
	CloseDB()
}

type dataBase struct {
	connection *gorm.DB
}

func NewFoodRepository() FoodRepository {
	db, err := gorm.Open("postgres", 'test.db')
	if err != nil {
		panic("Fail to connect database")
	}
	err = db.AutoMigrate(&entity.Food{}, &entity.Person{})
	if err != nil {
		panic("Fail in the migrate")
	}
	return &dataBase{
		connection: db,
	}
}

func (db *dataBase) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Fail to close database")
	}
}

func (db *dataBase) Save(food entity.Food) {
	db.connection.Create(&food)
}
func (db *dataBase) Update(food entity.Food) {
	db.connection.Save(&food)
}
func (db *dataBase) Delete(food entity.Food) {
	db.connection.Delete(&food)
}
func (db *dataBase) FindAll() []entity.Food {
	var foods []entity.Food
	db.connection.Set("gorm:auto_preload", true).Find(&foods)
	return foods
}
