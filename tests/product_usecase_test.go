package tests

import (
	"testing"

	"github.com/google/uuid"
	"github.com/vitormoschetta/go/src/domain/models"
	"github.com/vitormoschetta/go/src/domain/useCases"
	"github.com/vitormoschetta/go/tests/mock"
)

func Test_With_Product_Add_With_Valid_Data(t *testing.T) {
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

func Test_With_Product_Add_With_Invalid_Name(t *testing.T) {
	repository := mock.NewFakeProductRepository()
	useCase := useCases.NewProductUseCase(repository)

	request := models.Product{
		Name:  "",
		Price: 10,
	}

	response := useCase.Save(request)

	if len(response.Errors) == 0 {
		t.Errorf("Expected errors, got %v", response.Errors)
	}
}

func Test_With_Product_Add_With_Invalid_Price(t *testing.T) {
	repository := mock.NewFakeProductRepository()
	useCase := useCases.NewProductUseCase(repository)

	request := models.Product{
		Name:  "Product 1",
		Price: 0,
	}

	response := useCase.Save(request)

	if len(response.Errors) == 0 {
		t.Errorf("Expected errors, got %v", response.Errors)
	}
}

func Test_With_Product_Update_With_Valid_Data(t *testing.T) {
	repository := mock.NewFakeProductRepository()
	useCase := useCases.NewProductUseCase(repository)

	request := models.Product{
		ID:    uuid.New().String(),
		Name:  "Product 1",
		Price: 10,
	}

	response := useCase.Save(request)

	if len(response.Errors) > 0 {
		t.Errorf("Expected no errors, got %v", response.Errors)
	}

	request2 := models.Product{
		ID:    request.ID,
		Name:  "Product 2",
		Price: 20,
	}

	response = useCase.Update(request2)

	if len(response.Errors) > 0 {
		t.Errorf("Expected no errors, got %v", response.Errors)
	}
}

func Test_With_Product_Update_With_Invalid_ID(t *testing.T) {
	repository := mock.NewFakeProductRepository()
	useCase := useCases.NewProductUseCase(repository)

	request := models.Product{
		ID:    uuid.New().String(),
		Name:  "Product 1",
		Price: 10,
	}

	response := useCase.Save(request)

	if len(response.Errors) > 0 {
		t.Errorf("Expected no errors, got %v", response.Errors)
	}

	request2 := models.Product{
		ID:    uuid.New().String(),
		Name:  "Product 2",
		Price: 20,
	}

	response = useCase.Update(request2)

	if len(response.Errors) == 0 {
		t.Errorf("Expected errors, got %v", response.Errors)
	}
}

func Test_With_Product_Update_With_Invalid_Name(t *testing.T) {
	repository := mock.NewFakeProductRepository()
	useCase := useCases.NewProductUseCase(repository)

	request := models.Product{
		ID:    uuid.New().String(),
		Name:  "Product 1",
		Price: 10,
	}

	response := useCase.Save(request)

	if len(response.Errors) > 0 {
		t.Errorf("Expected no errors, got %v", response.Errors)
	}

	request2 := models.Product{
		ID:    request.ID,
		Name:  "",
		Price: 20,
	}

	response = useCase.Update(request2)

	if len(response.Errors) == 0 {
		t.Errorf("Expected errors, got %v", response.Errors)
	}
}

func Test_With_Product_Update_With_Invalid_Price(t *testing.T) {
	repository := mock.NewFakeProductRepository()
	useCase := useCases.NewProductUseCase(repository)

	request := models.Product{
		ID:    uuid.New().String(),
		Name:  "Product 1",
		Price: 10,
	}

	response := useCase.Save(request)

	if len(response.Errors) > 0 {
		t.Errorf("Expected no errors, got %v", response.Errors)
	}

	request2 := models.Product{
		ID:    request.ID,
		Name:  "Product 2",
		Price: 0,
	}

	response = useCase.Update(request2)

	if len(response.Errors) == 0 {
		t.Errorf("Expected errors, got %v", response.Errors)
	}
}
