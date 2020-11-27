package entity

type Food struct {
	Title      string `json:"title"`
	Ingredient string `json:"ingredient"`
	Price      string `json:"price"`
	IsExist    bool   `json:"is_exist"`
}
