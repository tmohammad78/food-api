package entity

import "time"

type Person struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `json:"first_name" binding:"required" gorm:"type:varchar(32)"`
	LastName  string `json:"last_name" binding:"required" gorm:"type:varchar(32)"`
	Age       int8   `json:"age" binding:"gte=1,lte=120"`
	Email     string `json:"email" binding:"required,email" gorm:"type:varchar(256)"`
}

type Food struct {
	ID         uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title      string    `json:"title" binding:"min=2,max=10" validate:"is-cool" gorm:"type:varchar(100)"`
	Ingredient string    `json:"ingredient" gorm:"type:varchar(100)"`
	Price      string    `json:"price" gorm:"type:int(100)"`
	IsExist    bool      `json:"is_exist" gorm:"type:bool"`
	Rate       int       `json:"rate" gorm:"type:int"`
	Author     Person    `json:"author" binding:"required" gorm:"foreignkey:PersonID"`
	PersonID   uint64    `json:"-" `
	CreatedAt  time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
