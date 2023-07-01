package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	categoryApplication "github.com/vitormoschetta/go/internal/application/category"
	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/domain/common"
	"github.com/vitormoschetta/go/internal/infra/api/requests"
	"github.com/vitormoschetta/go/internal/infra/api/responses"
	"github.com/vitormoschetta/go/pkg/utils"
)

type CategoryController struct {
	UseCase    *categoryApplication.CategoryUseCases
	Repository common.IRepository[category.Category]
}

func NewCategoryController(repository common.IRepository[category.Category], useCase *categoryApplication.CategoryUseCases) *CategoryController {
	return &CategoryController{
		UseCase:    useCase,
		Repository: repository,
	}
}

func (c *CategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var response responses.Response

	items, err := c.Repository.FindAll(ctx)
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

func (c *CategoryController) Get(w http.ResponseWriter, r *http.Request) {
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

func (c *CategoryController) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var response responses.Response

	var request requests.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = responses.ItemToResponse(request, "Bad request", ctx)
		json.NewEncoder(w).Encode(response)
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	input := categoryApplication.NewCreateCategoryInput(request.Name)
	output := c.UseCase.Create(ctx, *input)
	response = responses.OutputToResponse(output)
	BuildHttpStatusCode2(output, r.Method, w)
	json.NewEncoder(w).Encode(response)
}

func (c *CategoryController) Put(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var response responses.Response

	var request requests.UpdateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = responses.ItemToResponse(request, "Bad request", ctx)
		json.NewEncoder(w).Encode(response)
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	input := categoryApplication.NewUpdateCategoryInput(request.ID, request.Name)
	output := c.UseCase.Update(ctx, *input)
	response = responses.OutputToResponse(output)
	BuildHttpStatusCode2(output, r.Method, w)
	json.NewEncoder(w).Encode(response)
}

func (c *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	output := c.UseCase.Delete(ctx, id)
	response := responses.OutputToResponse(output)
	BuildHttpStatusCode2(output, r.Method, w)
	json.NewEncoder(w).Encode(response)
}
