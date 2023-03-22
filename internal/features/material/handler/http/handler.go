package materialHandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"price/internal/domain"
	"price/internal/entity"
	"price/internal/utils"
)

type Handler struct {
	UseCase domain.MaterialUsecase
}

func NewHandler(e *echo.Echo, useCase domain.MaterialUsecase) {
	handler := &Handler{
		UseCase: useCase,
	}

	e.GET("material/read", handler.GetMaterials)
	e.GET("material/update", handler.UpdateMaterials)
}

func (m *Handler) GetMaterials(e echo.Context) error {
	var formID entity.ID
	var formPage entity.PageStr
	var readRequest entity.ReadMaterialsRequest

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
	responses, err := m.UseCase.ReadMaterials(ctx, readRequest)
	if err != nil {
		return e.JSON(http.StatusNotFound, err.Error())
	}

	return e.JSON(http.StatusOK, responses)
}

func (m *Handler) UpdateMaterials(e echo.Context) error {
	var err error
	var items entity.UpdateMaterialRequest
	items.Sheet, err = utils.GetFromSheet()
	if err != nil {
		return e.JSON(http.StatusNotFound, err.Error())
	}

	ctx := e.Request().Context()
	if err := m.UseCase.UpdateMaterials(ctx, items); err != nil {
		return e.JSON(http.StatusNotFound, err.Error())
	}

	return e.JSON(http.StatusOK, "Updated.")
}
