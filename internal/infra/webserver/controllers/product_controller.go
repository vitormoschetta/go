package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	productApplication "github.com/vitormoschetta/go/internal/application/product"
	"github.com/vitormoschetta/go/internal/domain/product"
	"github.com/vitormoschetta/go/internal/share/utils"
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
	items, err := c.Repository.FindAll(ctx)
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
}

func (c *ProductController) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	item, err := c.Repository.FindByID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	}
	if item.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Product not found"}`))
		return
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
	vars := mux.Vars(r)
	categoryId := vars["category_id"]
	items, err := c.Repository.FindByCategory(ctx, categoryId)
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
	var input productApplication.CreateProductInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errOutput := utils.FormatErrOut(ctx, err)
		w.Write(errOutput)
		log.Println(string(errOutput))
		return
	}
	output := c.UseCase.Create(ctx, input)
	outputJSON, err := json.Marshal(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errOutput := utils.FormatErrOut(ctx, err)
		w.Write(errOutput)
		log.Println(string(errOutput))
		return
	}
	w.WriteHeader(output.Code)
	w.Write(outputJSON)
}

func (c *ProductController) Put(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input productApplication.UpdateProductInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errOutput := utils.FormatErrOut(ctx, err)
		w.Write(errOutput)
		log.Println(string(errOutput))
		return
	}
	output := c.UseCase.Update(ctx, input)
	outputJSON, err := json.Marshal(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errOutput := utils.FormatErrOut(ctx, err)
		w.Write(errOutput)
		log.Println(string(errOutput))
		return
	}
	w.WriteHeader(output.Code)
	w.Write(outputJSON)
}

func (c *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	output := c.UseCase.Delete(ctx, id)
	outputJSON, err := json.Marshal(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errOutput := utils.FormatErrOut(ctx, err)
		w.Write(errOutput)
		log.Println(string(errOutput))
		return
	}
	w.WriteHeader(output.Code)
	w.Write(outputJSON)
}

func (c *ProductController) PutPromotion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input productApplication.ApplyPromotionProductInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errOutput := utils.FormatErrOut(ctx, err)
		w.Write(errOutput)
		log.Println(string(errOutput))
		return
	}
	output := c.UseCase.ApplyPromotion(ctx, input)
	outputJSON, err := json.Marshal(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errOutput := utils.FormatErrOut(ctx, err)
		w.Write(errOutput)
		log.Println(string(errOutput))
		return
	}
	w.WriteHeader(output.Code)
	w.Write(outputJSON)
}

func (c *ProductController) PutPromotionbyCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input productApplication.ApplyPromotionProductByCategoryInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errOutput := utils.FormatErrOut(ctx, err)
		w.Write(errOutput)
		log.Println(string(errOutput))
	}
	output := c.UseCase.ApplyPromotionOnProductsByCategory(ctx, input)
	outputJSON, err := json.Marshal(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errOutput := utils.FormatErrOut(ctx, err)
		w.Write(errOutput)
		log.Println(string(errOutput))
	}
	w.WriteHeader(output.Code)
	w.Write(outputJSON)
	ctx.Done()
}
