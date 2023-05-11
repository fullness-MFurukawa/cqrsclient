package application_test

import (
	"client/application"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryList(t *testing.T) {
	service, err := application.NewQueryService()
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}

	categories, err := service.CagtegoryList()
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	log.Println(categories)
	assert.True(t, true)
}

func TestProductList(t *testing.T) {
	service, err := application.NewQueryService()
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}

	products, err := service.ProductList()
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	log.Println(products)
	assert.True(t, true)
}

func TestSearchCategoryById_OK(t *testing.T) {
	service, err := application.NewQueryService()
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	category, err := service.SearchCategoryById("b1524011-b6af-417e-8bf2-f449dd58b5c0")
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	log.Println(category)
	assert.True(t, true)
}

func TestSearchCategoryById_NG(t *testing.T) {
	service, err := application.NewQueryService()
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	category, err := service.SearchCategoryById("b1524011-b6af-417e-8bf2-f449dd58b5c1")
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	log.Println(category)
	assert.True(t, true)
}

func TestSearchProductById_OK(t *testing.T) {
	service, err := application.NewQueryService()
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	category, err := service.SearchProductById("38c6e236-90ca-48a2-b427-acb9d834b591")
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	log.Println(category)
	assert.True(t, true)
}

func TestSearchProductById_NG(t *testing.T) {
	service, err := application.NewQueryService()
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	product, err := service.SearchProductById("38c6e236-90ca-48a2-b427-acb9d834b590")
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	log.Println(product)
	assert.True(t, true)
}

func TestSearchProductByKeyword_OK(t *testing.T) {
	service, err := application.NewQueryService()
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	products, err := service.SearchProductByKeyword("ペン")
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	log.Println(products)
	assert.True(t, true)
}

func TestSearchProductByKeyword_NG(t *testing.T) {
	service, err := application.NewQueryService()
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	products, err := service.SearchProductByKeyword("おいうえお")
	if err != nil {
		log.Println(err.Error())
		assert.Error(t, err)
	}
	log.Println(products)
	assert.True(t, true)
}
