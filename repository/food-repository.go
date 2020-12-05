package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tmohammad78/food-api/entity"
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

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "godisbigest"
	dbname   = "restaurant"
)

func NewFoodRepository() FoodRepository {
	config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port,
		user, dbname, password)
	db, err := gorm.Open("postgres", config)
	if err != nil {
		panic("Fail to connect database")
	}
	db.AutoMigrate(&entity.Food{}, &entity.Person{})
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
	err := db.connection.Create(&food)
	if err != nil {
		fmt.Println(err)
	}
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
