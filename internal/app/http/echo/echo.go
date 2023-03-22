package echo

import (
	"github.com/labstack/echo/v4"
	"price/internal/domain"
	itemHandler "price/internal/features/item/handler/http"
	materialHandler "price/internal/features/material/handler/http"
	recipeHandler "price/internal/features/recipe/handler/http"
)

func New(s domain.Services) {
	e := echo.New()

	itemHandler.NewHandler(e, s.Item)
	materialHandler.NewHandler(e, s.Material)
	recipeHandler.NewHandler(e, s.Recipe)

	e.Logger.Fatal(e.Start(s.Port))
}

//func routeToService(routes []*echo.Route) (services []domain.CreateRouteRequest) {
//	for _, route := range routes {
//		f := strings.Split(route.Name, ".")
//		funcName := strings.Split(f[len(f)-1], "-")[0]
//		services = append(services, domain.CreateRouteRequest{
//			Name:     route.Method + "-" + route.Path,
//			Code:     f[len(f)-1],
//			Path:     route.Path,
//			Function: funcName,
//			Method:   route.Method,
//			IsActive: true,
//		})
//	}
//
//	return
//}
