package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	categoryApplication "github.com/vitormoschetta/go/internal/application/category"
	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/domain/common"
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
	w.Header().Set("Content-Type", "application/json")
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
	ctx.Done()
}

func (c *CategoryController) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	item, err := c.Repository.FindByID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	}
	responseItemJSON, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseItemJSON)
	ctx.Done()
}

func (c *CategoryController) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	var input categoryApplication.CreateCategoryInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}
	response, statusCode := c.UseCase.Create(ctx, input)
	w.WriteHeader(statusCode)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	}
	w.Write(responseJSON)
	ctx.Done()
}

func (c *CategoryController) Put(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	var input categoryApplication.UpdateCategoryInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}
	response, statusCode := c.UseCase.Update(ctx, input)
	w.WriteHeader(statusCode)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	}
	w.Write(responseJSON)
	ctx.Done()
}

func (c *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	response, statusCode := c.UseCase.Delete(ctx, id)
	w.WriteHeader(statusCode)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	}
	w.Write(responseJSON)
	ctx.Done()
}
