package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"latihan-golang-unit-test/entity"
	"latihan-golang-unit-test/repository"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, err := categoryService.Get("1")
	assert.Nil(t, category, nil)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "Laptop"}

	categoryRepository.Mock.On("FindById", category.Id).Return(category)
	result, err := categoryService.Get(category.Id)
	assert.Nil(t, err)
	assert.NotNil(t, category.Id, result.Id)
	assert.NotNil(t, category.Name, result.Name)
}
