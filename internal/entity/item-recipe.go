package entity

import (
	"time"
)

type ItemRecipe struct {
	ID        uint      `bun:"id,pk,autoincrement"`
	CreatedAt time.Time `bun:",nullzero,default:current_timestamp"`
	UpdatedAt time.Time
	DeletedAt time.Time `bun:",nullzero,default:current_timestamp"`
	ItemID    uint
	RecipeID  uint
}

type ItemRecipeResponse struct {
	ID        uint
	Item      ItemResponse
	Recipe    RecipeResponse
	UpdatedAt time.Time
}

type CreateItemRecipeRequest struct {
	ItemID   uint
	RecipeID uint
}

type ReadItemRecipeRequest struct {
	Usual    *UsualRequest
	ItemID   ID
	RecipeID ID
}

type UpdateItemRecipeRequest struct {
	ID        uint
	ItemID    uint
	RecipeID  uint
	UpdatedAt time.Time
}

type DeleteItemRecipeRequest struct {
	ID        uint
	DeletedAt time.Time
}
