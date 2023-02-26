package tests

import (
	"testing"

	"github.com/vitormoschetta/go/models"
	"github.com/vitormoschetta/go/tests/mock"
	"github.com/vitormoschetta/go/useCases"
)

func TestProductRepository(t *testing.T) {
	repository := mock.NewFakeProductRepository()
	useCase := useCases.NewProductUseCase(repository)

	request := models.Product{
		Name:  "Product 1",
		Price: 10,
	}

	response := useCase.Save(request)

	if len(response.Errors) > 0 {
		t.Errorf("Expected no errors, got %v", response.Errors)
	}
}
