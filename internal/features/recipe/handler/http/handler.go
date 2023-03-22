package recipeHandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"price/internal/domain"
	"price/internal/entity"
)

type Handler struct {
	UseCase domain.RecipeUsecase
}

func NewHandler(e *echo.Echo, useCase domain.RecipeUsecase) {
	handler := &Handler{
		UseCase: useCase,
	}

	e.POST("recipe/create", handler.CreateRecipe)
	e.GET("recipe/read", handler.GetRecipes)
	e.PATCH("recipe/update", handler.UpdateRecipe)
}

func (m *Handler) CreateRecipe(e echo.Context) error {
	var request entity.CreateRecipeRequest
	if err := e.Bind(&request); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := e.Request().Context()
	if err := m.UseCase.CreateRecipe(ctx, request); err != nil {
		return e.JSON(http.StatusNotFound, err.Error())
	}

	return e.JSON(http.StatusCreated, "created.")
}

func (m *Handler) GetRecipes(e echo.Context) error {
	var formID entity.ID
	var formPage entity.PageStr
	var readRequest entity.ReadRecipeRequest

	formID.Str = e.QueryParam("id")
	formPage.PageNumber = e.QueryParam("page")
	formPage.PerPage = e.QueryParam("perpage")

	request := &entity.UsualRequest{
		ID:          formID,
		PageRequest: formPage,
		Name:        e.QueryParam("name"),
	}

	readRequest.Usual = request

	ctx := e.Request().Context()
	responses, err := m.UseCase.ReadRecipe(ctx, readRequest)
	if err != nil {
		return e.JSON(http.StatusNotFound, err.Error())
	}

	return e.JSON(http.StatusOK, responses)
}

func (m *Handler) UpdateRecipe(e echo.Context) error {
	var request entity.UpdateRecipeRequest
	if err := e.Bind(&request); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := e.Request().Context()
	if err := m.UseCase.UpdateRecipe(ctx, request); err != nil {
		return e.JSON(http.StatusNotFound, err.Error())
	}

	return e.JSON(http.StatusOK, "Updated.")
}
