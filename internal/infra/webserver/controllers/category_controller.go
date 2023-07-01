package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	categoryApplication "github.com/vitormoschetta/go/internal/application/category"
	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/domain/common"
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
	items, err := c.Repository.FindAll(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Internal error"))
		// json.NewEncoder(w).Encode(utils.FormatErrOutWithMessage2(ctx, "Internal error"))
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	json.NewEncoder(w).Encode(items)
}

func (c *CategoryController) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	item, err := c.Repository.FindByID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// json.NewEncoder(w).Encode(utils.FormatErrOutWithMessage2(ctx, "Internal error"))
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Internal error"))
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	if item.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		// json.NewEncoder(w).Encode(utils.FormatErrOutWithMessage2(ctx, "Not found"))
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Not found"))
		log.Print(utils.BuildLoggerWithErr3(ctx, "Not found", utils.GetCallingPackage()))
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (c *CategoryController) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input categoryApplication.CreateCategoryInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Bad request"))
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	output := c.UseCase.Create(ctx, input)
	response := OutputToResponse(output)
	BuildHttpStatusCode2(output, VerbTypePost, &w)
	json.NewEncoder(w).Encode(response)
}

func (c *CategoryController) Put(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input categoryApplication.UpdateCategoryInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Bad request"))
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	output := c.UseCase.Update(ctx, input)
	response := OutputToResponse(output)
	BuildHttpStatusCode2(output, VerbTypePost, &w)
	json.NewEncoder(w).Encode(response)
}

func (c *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	output := c.UseCase.Delete(ctx, id)
	response := OutputToResponse(output)
	BuildHttpStatusCode2(output, VerbTypePost, &w)
	json.NewEncoder(w).Encode(response)
}
