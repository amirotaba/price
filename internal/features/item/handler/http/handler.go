package itemHandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"price/internal/domain"
	"price/internal/entity"
)

type Handler struct {
	UseCase domain.ItemUsecase
}

func NewHandler(e *echo.Echo, useCase domain.ItemUsecase) {
	handler := &Handler{
		UseCase: useCase,
	}

	e.POST("item/create", handler.CreateItem)
	e.GET("item/read", handler.GetItems)
	e.PATCH("item/update", handler.UpdateItem)
}

func (m *Handler) CreateItem(e echo.Context) error {
	var request entity.CreateItemRequest
	if err := e.Bind(&request); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := e.Request().Context()
	if err := m.UseCase.CreateItem(ctx, request); err != nil {
		return e.JSON(http.StatusNotFound, err.Error())
	}

	return e.JSON(http.StatusCreated, "created.")
}

func (m *Handler) GetItems(e echo.Context) error {
	var formID entity.ID
	var formPage entity.PageStr
	var readRequest entity.ReadItemRequest

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
	responses, err := m.UseCase.ReadItem(ctx, readRequest)
	if err != nil {
		return e.JSON(http.StatusNotFound, err.Error())
	}

	return e.JSON(http.StatusOK, responses)
}

func (m *Handler) UpdateItem(e echo.Context) error {
	var request entity.UpdateItemRequest
	if err := e.Bind(&request); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := e.Request().Context()
	if err := m.UseCase.UpdateItem(ctx, request); err != nil {
		return e.JSON(http.StatusNotFound, err.Error())
	}

	return e.JSON(http.StatusOK, "Updated.")
}
