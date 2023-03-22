package domain

import (
	"golang.org/x/net/context"
	"price/internal/entity"
)

type ItemRecipeRepository interface {
	CreateItemRecipe(ctx context.Context, request *entity.CreateItemRecipeRequest) error
	ReadItemRecipe(ctx context.Context, request *ReadRequest) ([]entity.ItemRecipe, error)
	UpdateItemRecipe(ctx context.Context, request *entity.UpdateItemRecipeRequest) error
	DeleteItemRecipe(ctx context.Context, request *entity.DeleteItemRecipeRequest) error
}
