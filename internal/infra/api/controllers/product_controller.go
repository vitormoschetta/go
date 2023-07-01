package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	productApplication "github.com/vitormoschetta/go/internal/application/product"
	"github.com/vitormoschetta/go/internal/domain/product"
	"github.com/vitormoschetta/go/internal/infra/api/responses"
	"github.com/vitormoschetta/go/pkg/utils"
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

func (c *ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	items, err := c.Repository.FindAll(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Internal error"))
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return
	}
	json.NewEncoder(w).Encode(items)
}

func (c *ProductController) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	item, err := c.Repository.FindByID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Internal error"))
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return
	}
	if item.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Not found"))
		log.Print(utils.BuildLogger(ctx, "Not found") + " - " + utils.GetCallingPackage())
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (c *ProductController) GetByCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	categoryId := vars["category_id"]
	items, err := c.Repository.FindByCategory(ctx, categoryId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Internal error"))
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return
	}
	json.NewEncoder(w).Encode(items)
}

func (c *ProductController) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input productApplication.CreateProductInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Bad request"))
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return
	}
	output := c.UseCase.Create(ctx, input)
	response := responses.OutputToResponse(output)
	BuildHttpStatusCode2(output, r.Method, w)
	json.NewEncoder(w).Encode(response)
}

func (c *ProductController) Put(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input productApplication.UpdateProductInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Bad request"))
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return
	}
	output := c.UseCase.Update(ctx, input)
	response := responses.OutputToResponse(output)
	BuildHttpStatusCode2(output, r.Method, w)
	json.NewEncoder(w).Encode(response)
}

func (c *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	output := c.UseCase.Delete(ctx, id)
	response := responses.OutputToResponse(output)
	BuildHttpStatusCode2(output, r.Method, w)
	json.NewEncoder(w).Encode(response)
}

func (c *ProductController) PutPromotion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input productApplication.ApplyPromotionProductInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Bad request"))
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return
	}
	output := c.UseCase.ApplyPromotion(ctx, input)
	response := responses.OutputToResponse(output)
	BuildHttpStatusCode2(output, r.Method, w)
	json.NewEncoder(w).Encode(response)
}

func (c *ProductController) PutPromotionbyCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input productApplication.ApplyPromotionProductByCategoryInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Bad request"))
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return
	}
	output := c.UseCase.ApplyPromotionOnProductsByCategory(ctx, input)
	response := responses.OutputToResponse(output)
	BuildHttpStatusCode2(output, r.Method, w)
	json.NewEncoder(w).Encode(response)
}
