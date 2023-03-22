package domain

import (
	"golang.org/x/net/context"
	"price/internal/entity"
)

type RecipeUsecase interface {
	CreateRecipe(ctx context.Context, request entity.CreateRecipeRequest) error
	ReadRecipe(ctx context.Context, request entity.ReadRecipeRequest) ([]entity.RecipeResponse, error)
	UpdateRecipe(ctx context.Context, request entity.UpdateRecipeRequest) error
}

type RecipeRepository interface {
	CreateRecipe(ctx context.Context, request entity.CreateRecipeRequest) error
	ReadRecipe(ctx context.Context, request *ReadRequest) ([]entity.Recipe, error)
	UpdateRecipe(ctx context.Context, request entity.UpdateRecipeRequest) error
}
