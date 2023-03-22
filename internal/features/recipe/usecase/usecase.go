package recipeUsecase

import (
	"errors"
	"github.com/jinzhu/copier"
	"golang.org/x/net/context"
	"price/internal/domain"
	"price/internal/entity"
	"price/internal/utils"
	"time"
)

type usecase struct {
	Recipe   domain.RecipeRepository
	Material domain.MaterialRepository
}

func NewUseCase(r domain.Repositories) domain.RecipeUsecase {
	return &usecase{
		Recipe:   r.Recipe,
		Material: r.Material,
	}
}

func (u *usecase) CreateRecipe(ctx context.Context, request entity.CreateRecipeRequest) error {
	var filters []domain.Filter
	filter := utils.AddFilter("id", uint(request.MaterialID), "equal")
	filters = append(filters, filter)

	readRequest := &domain.ReadRequest{
		Filters: filters,
	}

	exist, _ := u.Material.ReadMaterial(ctx, readRequest)
	if exist == nil {
		return errors.New("material doesn't found")
	}

	if err := u.Recipe.CreateRecipe(ctx, request); err != nil {
		return err
	}
	return nil
}

func (u *usecase) ReadRecipe(ctx context.Context, request entity.ReadRecipeRequest) ([]entity.RecipeResponse, error) {
	var response entity.RecipeResponse
	var responses []entity.RecipeResponse
	convRequest, err := utils.ConvertRequest(request.Usual)
	if err != nil {
		return nil, err
	}

	recipes, err := u.Recipe.ReadRecipe(ctx, convRequest)

	for _, recipe := range recipes {
		if err := copier.Copy(&response, &recipe); err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (u *usecase) UpdateRecipe(ctx context.Context, request entity.UpdateRecipeRequest) error {
	var filters []domain.Filter
	filter := utils.AddFilter("id", uint(request.MaterialID), "equal")
	filters = append(filters, filter)

	readRequest := &domain.ReadRequest{
		Filters: filters,
	}

	exist, _ := u.Material.ReadMaterial(ctx, readRequest)
	if exist == nil {
		return errors.New("material doesn't found")
	}

	request.UpdatedAt = time.Now()

	if err := u.Recipe.UpdateRecipe(ctx, request); err != nil {
		return err
	}
	return nil
}
