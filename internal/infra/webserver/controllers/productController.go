package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	productApplication "github.com/vitormoschetta/go/internal/application/product"
	"github.com/vitormoschetta/go/internal/domain/product"
)

type ProductController struct {
	UseCase    *productApplication.ProductUseCase
	Repository product.IProductRepository
}

func NewProductController(repository product.IProductRepository, useCase *productApplication.ProductUseCase) *ProductController {
	return &ProductController{
		UseCase:    useCase,
		Repository: repository,
	}
}

// Show Products godoc
//
// @Summary      Show all products
// @Description  Get all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Product
// @Router       /products [get]
func (c *ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	items, err := c.Repository.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	}
	responseItemsJSON, err := json.Marshal(items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseItemsJSON)
	ctx.Done()
}

func (c *ProductController) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	item, err := c.Repository.FindByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "Product not found"}`))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	}
	responseItemsJSON, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseItemsJSON)
	ctx.Done()
}

func (c *ProductController) GetByCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	categoryId := vars["category_id"]
	items, err := c.Repository.FindByCategory(categoryId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	}
	responseItemsJSON, err := json.Marshal(items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseItemsJSON)
	ctx.Done()
}

func (c *ProductController) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	var input productApplication.CreateProductInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}
	response, statusCode := c.UseCase.Create(input)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	}
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
	ctx.Done()
}

func (c *ProductController) Put(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	var input productApplication.UpdateProductInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}
	response, statusCode := c.UseCase.Update(input)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	}
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
	ctx.Done()
}

func (c *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	response, statusCode := c.UseCase.Delete(id)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	}
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
	ctx.Done()
}

func (c *ProductController) PutPromotion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	var input productApplication.ApplyPromotionProductInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}
	response, statusCode := c.UseCase.ApplyPromotion(input)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	}
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
	ctx.Done()
}

func (c *ProductController) PutPromotionbyCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input productApplication.ApplyPromotionProductByCategoryInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}
	response, statusCode := c.UseCase.ApplyPromotionOnProductsByCategory(input)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	}
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
	ctx.Done()
}
