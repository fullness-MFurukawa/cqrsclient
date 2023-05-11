package dto

type ProductDTO struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Price      int32  `json:"price"`
	CategoryId string `json:"categoryId"`
}
