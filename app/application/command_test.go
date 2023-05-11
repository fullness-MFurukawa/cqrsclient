package application_test

import (
	"client/application"
	"client/presentation/dto"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCategory(t *testing.T) {
	service, err := application.NewCommandService()
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	dto := dto.CategoryDTO{Id: "", Name: "食品"}
	result, err := service.CreateCategory(dto)
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	log.Println(*result)
	assert.True(t, true)
}

func TestUpdateCategory(t *testing.T) {
	service, err := application.NewCommandService()
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	dto := dto.CategoryDTO{Id: "f2ac0895-b23b-4379-b7aa-ba076bb9933c", Name: "食品!!!!!!!!"}
	result, err := service.UpdateCategory(dto)
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	log.Println(*result)
	assert.True(t, true)
}

func TestDeleteategory(t *testing.T) {
	service, err := application.NewCommandService()
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	dto := dto.CategoryDTO{Id: "f2ac0895-b23b-4379-b7aa-ba076bb9933c"}
	result, err := service.DeleteCategory(dto)
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	log.Println(*result)
	assert.True(t, true)
}

func TestCreateProduct(t *testing.T) {
	service, err := application.NewCommandService()
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	dto := dto.ProductDTO{Name: "テスト商品", Price: int32(500), CategoryId: "762bd1ea-9700-4bab-a28d-6cbebf20ddc2"}
	result, err := service.CreateProduct(dto)
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	log.Println(*result)
	assert.True(t, true)
}

func TestUpdateProduct(t *testing.T) {
	service, err := application.NewCommandService()
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	dto := dto.ProductDTO{Id: "9f9aa728-bef6-4dda-ae1f-98307b0f0573", Name: "テスト商品@@@@", Price: int32(500), CategoryId: "762bd1ea-9700-4bab-a28d-6cbebf20ddc2"}
	result, err := service.UpdateProduct(dto)
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	log.Println(*result)
	assert.True(t, true)
}

func TestDeleteProduct(t *testing.T) {
	service, err := application.NewCommandService()
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	dto := dto.ProductDTO{Id: "9f9aa728-bef6-4dda-ae1f-98307b0f0573"}
	result, err := service.DeleteProduct(dto)
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	log.Println(*result)
	assert.True(t, true)
}
