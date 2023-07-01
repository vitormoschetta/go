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

// Show Categories godoc
//
// @Summary      Show all categories
// @Description  Get all categories
// @Tags         categories
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Category
// @Router       /categories [get]
func (c *CategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	items, err := c.Repository.FindAll(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Internal error"))
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	responseItemsJSON, err := json.Marshal(items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Internal error"))
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(responseItemsJSON)
}

func (c *CategoryController) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	item, err := c.Repository.FindByID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Internal error"))
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	if item.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Not found"))
		log.Print(utils.BuildLoggerWithErr3(ctx, "Not found", utils.GetCallingPackage()))
		return
	}
	responseItemJSON, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Internal error"))
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(responseItemJSON)
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
	outputJSON, err := json.Marshal(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Internal error"))
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	w.WriteHeader(BuildHttpStatusCode(output.Code, VerbTypePost))
	_, _ = w.Write(outputJSON)
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
	outputJSON, err := json.Marshal(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Internal error"))
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	w.WriteHeader(BuildHttpStatusCode(output.Code, VerbTypePut))
	_, _ = w.Write(outputJSON)
}

func (c *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	output := c.UseCase.Delete(ctx, id)
	outputJSON, err := json.Marshal(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(utils.FormatErrOutWithMessage(ctx, "Internal error"))
		log.Print(utils.BuildLoggerWithErr2(ctx, err, utils.GetCallingPackage()))
		return
	}
	w.WriteHeader(BuildHttpStatusCode(output.Code, VerbTypeDelete))
	_, _ = w.Write(outputJSON)
}
