package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	productApplication "github.com/vitormoschetta/go/internal/application/product"
	"github.com/vitormoschetta/go/internal/domain/product"
	"github.com/vitormoschetta/go/internal/infra/api/requests"
	"github.com/vitormoschetta/go/internal/infra/api/responses"
	"github.com/vitormoschetta/go/pkg/pagination"
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
	var response responses.Response
	var pagination pagination.Pagination
	var err error

	queryParams := r.URL.Query()
	pagination.CurrentPage, err = strconv.Atoi(queryParams.Get("page"))
	if err != nil {
		pagination.CurrentPage = DefaultCurrentPage
	}
	pagination.PageSize, err = strconv.Atoi(queryParams.Get("page_size"))
	if err != nil {
		pagination.PageSize = DefaultPageSize
	}

	items, err := c.Repository.FindAll(ctx, &pagination)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response = responses.ItemToResponse(items, "Internal error", ctx)
		json.NewEncoder(w).Encode(response)
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	response = responses.ItemToResponseWithPagination(items, "", ctx, &pagination)
	json.NewEncoder(w).Encode(response)
}

func (c *ProductController) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var response responses.Response

	vars := mux.Vars(r)
	id := vars["id"]
	item, err := c.Repository.FindByID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response = responses.ItemToResponse(item, "Internal error", ctx)
		json.NewEncoder(w).Encode(response)
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	if item.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		response = responses.ItemToResponse(item, "Not found", ctx)
		json.NewEncoder(w).Encode(response)
		log.Print(utils.BuildLoggerWithErr3(ctx, "Not found", utils.GetCallingPackage()))
		return
	}
	response = responses.ItemToResponse(item, "", ctx)
	json.NewEncoder(w).Encode(response)
}

func (c *ProductController) GetByCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var response responses.Response

	vars := mux.Vars(r)
	categoryId := vars["category_id"]
	items, err := c.Repository.FindByCategory(ctx, categoryId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response = responses.ItemToResponse(items, "Internal error", ctx)
		json.NewEncoder(w).Encode(response)
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	response = responses.ItemToResponse(items, "", ctx)
	json.NewEncoder(w).Encode(response)
}

func (c *ProductController) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var response responses.Response

	var request requests.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = responses.ItemToResponse(request, "Bad request", ctx)
		json.NewEncoder(w).Encode(response)
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	input := productApplication.NewCreateProductInput(request.Name, request.Price, request.CategoryID)
	output := c.UseCase.Create(ctx, *input)
	response = responses.OutputToResponse(output)
	BuildHttpStatusCode2(output, r.Method, w)
	json.NewEncoder(w).Encode(response)
}

func (c *ProductController) Put(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var response responses.Response

	var request requests.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = responses.ItemToResponse(request, "Bad request", ctx)
		json.NewEncoder(w).Encode(response)
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	input := productApplication.NewUpdateProductInput(request.ID, request.Name, request.Price, request.CategoryID)
	output := c.UseCase.Update(ctx, *input)
	response = responses.OutputToResponse(output)
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
	var response responses.Response

	var request requests.ApplyPromotionProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = responses.ItemToResponse(request, "Bad request", ctx)
		json.NewEncoder(w).Encode(response)
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	input := productApplication.NewApplyPromotionProductInput(request.ProductId, request.Percentage)
	output := c.UseCase.ApplyPromotion(ctx, *input)
	response = responses.OutputToResponse(output)
	BuildHttpStatusCode2(output, r.Method, w)
	json.NewEncoder(w).Encode(response)
}

func (c *ProductController) PutPromotionbyCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var response responses.Response

	var request requests.ApplyPromotionProductByCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = responses.ItemToResponse(request, "Bad request", ctx)
		json.NewEncoder(w).Encode(response)
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	input := productApplication.NewApplyPromotionProductByCategoryInput(request.CategoryId, request.Percentage)
	output := c.UseCase.ApplyPromotionOnProductsByCategory(ctx, *input)
	response = responses.OutputToResponse(output)
	BuildHttpStatusCode2(output, r.Method, w)
	json.NewEncoder(w).Encode(response)
}
